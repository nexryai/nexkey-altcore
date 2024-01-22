package migration

import (
	"database/sql"
	"github.com/lopezator/migrator"
)

var migrationsToAltcore = migrator.Migrations(
	&migrator.Migration{
		Name: "Drop channel",
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec("DROP INDEX \"IDX_note_on_channelId_and_id_desc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_6a57f051d82c6d4036c141e107\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_29e8c1d579af54d4232939f994\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_89a29c9237b8c3b6b3cbb4cb30\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_25b1dd384bec391b07b74b861c\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" DROP COLUMN \"noteChannelId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" DROP COLUMN \"isMentioned\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" ADD \"readCursor\" TIMESTAMP WITH TIME ZONE NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" DROP COLUMN \"readCursor\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_note_pining\" DROP CONSTRAINT \"FK_10b19ef67d297ea9de325cd4502\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_note_pining\" DROP CONSTRAINT \"FK_8125f950afd3093acb10d2db8a8\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" DROP CONSTRAINT \"FK_6d8084ec9496e7334a4602707e1\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" DROP CONSTRAINT \"FK_0e43068c3f92cab197c3d3cd86e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP CONSTRAINT \"FK_f22169eb10657bded6d875ac8f9\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel\" DROP CONSTRAINT \"FK_999da2bcc7efadbfe0e92d3bc19\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DALTER TABLE \"channel\" DROP CONSTRAINT \"FK_823bae55bd81b3be6e05cff4383\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP COLUMN \"channelId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f36fed37d6d4cdcc68c803cd9c\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_8125f950afd3093acb10d2db8a\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"channel_note_pining\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_2e230dd45a10e671d781d99f3e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_6d8084ec9496e7334a4602707e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0e43068c3f92cab197c3d3cd86\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_11e71f2511589dcc8a4d3214f9\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"channel_following\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_094b86cd36bb805d1aa1e8cc9a\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0f58c11241e649d2a638a8de94\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_823bae55bd81b3be6e05cff438\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_29ef80c6f13bcea998447fce43\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_71cb7b435b7c0d4843317e7e16\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"channel\""); err != nil {
				return err
			}

			return nil
		},
	},
	&migrator.Migration{
		Name: "Drop gallery",
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec("ALTER TABLE \"gallery_like\" DROP CONSTRAINT \"FK_b1cb568bfe569e47b7051699fc8\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"gallery_like\" DROP CONSTRAINT \"FK_8fd5215095473061855ceb948cf\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"gallery_post\" DROP CONSTRAINT \"FK_985b836dddd8615e432d7043ddb\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_df1b5f4099e99fb0bc5eae53b6\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_8fd5215095473061855ceb948c\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"gallery_like\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_05cca34b985d1b8edc1d1e28df\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_1a165c68a49d08f11caffbd206\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f2d744d9a14d0dfb8b96cb7fc5\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_3ca50563facd913c425e7a89ee\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_985b836dddd8615e432d7043dd\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f631d37835adb04792e361807c\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_8f1a239bd077c8864a20c62c2c\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"gallery_post\""); err != nil {
				return err
			}
			return nil
		},
	},
	&migrator.Migration{
		Name: "Drop unused tables",
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec("ALTER TABLE \"reversi_matching\" DROP CONSTRAINT \"FK_e247b23a3c9b45f89ec1299d066\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_matching\" DROP CONSTRAINT \"FK_3b25402709dd9882048c2bbade0\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_game\" DROP CONSTRAINT \"FK_6649a4e8c5d5cf32fb03b5da9f6\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_game\" DROP CONSTRAINT \"FK_f7467510c60a45ce5aca6292743\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"reversi_matching\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CDROP TABLE \"reversi_game\""); err != nil {
				return err
			}
			return nil
		},
	},
	&migrator.Migration{
		Name: "Fix tables types",
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec("ALTER TABLE user_profile DROP COLUMN \"mutedWords\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE user_profile DROP COLUMN \"mutedInstances\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE user_profile ADD COLUMN \"mutedWords\" VARCHAR(255)[] DEFAULT '{}'::VARCHAR(255)[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE user_profile ADD COLUMN \"mutedInstances\" VARCHAR(255)[] DEFAULT '{}'::VARCHAR(255)[]"); err != nil {
				return err
			}
			return nil
		},
	})
