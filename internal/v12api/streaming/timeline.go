package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/db/kv"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"log"
	"time"
)

var ctx = context.Background()

type connectMessageBody struct {
	Channel string `json:"channel"`
	ID      string `json:"id"`
}

type connectMessage struct {
	Type string             `json:"type"`
	Body connectMessageBody `json:"body"`
}

type notificationMessage struct {
	Type string              `json:"type"`
	Body schema.Notification `json:"body"`
}

type channels struct {
	subscribeSystemStats bool
	subscribeNotify      bool
	subscribeTimeline    enum.TimelineType
}

func shouldAddToTimeline(followees *[]string, tl enum.TimelineType, note *schema.Note) bool {
	switch tl {
	case enum.HomeTimeline:
		// フォローしているユーザーのノートのみを抽出
		for _, followee := range *followees {
			if note.UserId == followee {
				return true
			}
		}

		return false
	}

	return false
}

func HandleStreamingApi(ctx *websocket.Conn) {
	ctx.Query("i")

	token, err := jwt.Parse(ctx.Query("i"), func(token *jwt.Token) (interface{}, error) {
		return []byte("DONOTUSEINPRODUCTION"), nil // Verify signing key
	})

	if err != nil || !token.Valid {
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	// 復号化できてるはずなのにuserIdがない場合panic
	if userId == "" {
		panic("JWT token was successfully decrypted but claims are invalid")
	}

	log.Println(userId)

	// ストリーミング用のRedisの接続確立
	//redisClient := kv.ConnectToRedis()

	// Subscribe用のチャネル
	//subCh := redisClient.Subscribe(ctx, fmt.Sprintf("%s/streaming", config.Host)).Channel()

	// サブスクライブするチャンネルを格納する共有メモリ
	listenChannels := &channels{
		subscribeNotify: true,
	}

	// タイムラインのためにキャッシュする情報
	followService := baselib.FollowService{
		Type:   enum.Followees,
		UserId: userId,
	}

	followees, err := followService.FindAll()
	if err != nil {
		logger.ErrorWithDetail("Error getting followees:", err)
		return
	}

	// サーバーメトリクスやジョブキューステータスを定期的に送信するスレッド
	go func(ch *channels) {
		for {
			time.Sleep(2 * time.Second)
			if ch.subscribeSystemStats {
				if err := ctx.WriteMessage(websocket.TextMessage, []byte("{\"type\":\"channel\",\"body\":{\"id\":\"3\",\"type\":\"stats\",\"body\":{\"cpu\":0.02,\"mem\":{\"used\":4716752896,\"active\":4243943424},\"net\":{\"rx\":1319.5,\"tx\":1400.6},\"fs\":{\"r\":0,\"w\":0}}}}")); err != nil {
					fmt.Println("Write error:", err)
					return
				}
			}
		}
	}(listenChannels)

	// タイムライン
	go func(ch *channels) {
		redisCtx := context.Background()

		redisClient := kv.ConnectToRedis()
		subscriber := redisClient.Subscribe(redisCtx, fmt.Sprintf("%s/streaming/timeline", config.Host))
		defer subscriber.Close()

		for {
			msg, err := subscriber.ReceiveMessage(redisCtx)
			if err != nil {
				logger.ErrorWithDetail("Error receiving message:", err)
				return
			}
			fmt.Println("Received:", msg.Payload)

			var note schema.Note
			err = json.Unmarshal([]byte(msg.Payload), &note)
			if err != nil {
				logger.ErrorWithDetail("Error unmarshalling notification:", err)
				return
			}

			// タイムラインを構築
			if shouldAddToTimeline(&followees, ch.subscribeTimeline, &note) {
				n := schema.Note{
					Id:         note.Id,
					UserId:     note.UserId,
					Visibility: note.Visibility,
					CreatedAt:  note.CreatedAt,
					Text:       note.Text,
				}

				jsonMsg, err := json.Marshal(n)
				if err != nil {
					logger.ErrorWithDetail("Error marshalling note:", err)
					return
				}

				err = ctx.WriteMessage(websocket.TextMessage, jsonMsg)
				if err != nil {
					fmt.Println("Write error:", err)
					return
				}

			} else {
				continue
			}
		}
	}(listenChannels)

	// 通知
	go func(ch *channels) {
		redisCtx := context.Background()

		redisClient := kv.ConnectToRedis()
		subscriber := redisClient.Subscribe(redisCtx, fmt.Sprintf("%s/streaming/notify/%s", config.Host, userId))
		defer subscriber.Close()

		for {
			msg, err := subscriber.ReceiveMessage(redisCtx)
			if ch.subscribeNotify {
				if err != nil {
					logger.ErrorWithDetail("Error receiving message:", err)
					return
				}
				fmt.Println("Received:", msg.Payload)

				var notification entities.Notification
				err = json.Unmarshal([]byte(msg.Payload), &notification)
				if err != nil {
					logger.ErrorWithDetail("Error unmarshalling notification:", err)
					return
				}

				if notification.NotifieeId == userId {
					n := notificationMessage{
						Type: "channel",
						Body: schema.Notification{
							Id:        notification.Id,
							CreatedAt: notification.CreatedAt,
							Type:      notification.Type,
							IsRead:    notification.IsRead,
							Note:      schema.Note{},
							User:      schema.User{},
						},
					}

					if notification.Type == "reaction" {
						n.Body.Reaction = notification.Reaction
					}

					jsonMsg, err := json.Marshal(n)
					if err != nil {
						logger.ErrorWithDetail("Error marshalling notification:", err)
						return
					}

					err = ctx.WriteMessage(websocket.TextMessage, jsonMsg)
					if err != nil {
						fmt.Println("Write error:", err)
						return
					}
				}
			}
		}
	}(listenChannels)

	// クライアントからのメッセージを待機
	var (
		msg []byte
	)
	for {
		if _, msg, err = ctx.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		// サブスクライブするチャンネルを変える
		var cntMsg connectMessage
		err := json.Unmarshal(msg, &cntMsg)
		if err != nil {
			fmt.Println("JSONパースエラー:", err)
		} else {
			logger.Info(cntMsg.Body.Channel)
			if cntMsg.Body.Channel == "serverStats" {
				listenChannels.subscribeSystemStats = true
			} else if cntMsg.Body.Channel == "homeTimeline" {
				listenChannels.subscribeTimeline = enum.HomeTimeline
			}
		}

	}
}
