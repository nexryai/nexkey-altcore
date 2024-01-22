package migration

import (
	"database/sql"
	"github.com/lopezator/migrator"
)

var migrationsToInitMisskey = migrator.Migrations(
	&migrator.Migration{
		Name: "Init DB",
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec("CREATE TYPE \"log_level_enum\" AS ENUM('error', 'warning', 'info', 'success', 'debug')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"log\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"domain\" character varying(64) array NOT NULL DEFAULT '{}'::varchar[], \"level\" \"log_level_enum\" NOT NULL, \"worker\" character varying(8) NOT NULL, \"machine\" character varying(128) NOT NULL, \"message\" character varying(1024) NOT NULL, \"data\" jsonb NOT NULL DEFAULT '{}', CONSTRAINT \"PK_350604cbdf991d5930d9e618fbd\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8e4eb51a35d81b64dda28eed0a\" ON \"log\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8cb40cfc8f3c28261e6f887b03\" ON \"log\" (\"domain\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_584b536b49e53ac81beb39a177\" ON \"log\" (\"level\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"drive_folder\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"name\" character varying(128) NOT NULL, \"userId\" character varying(32), \"parentId\" character varying(32), CONSTRAINT \"PK_7a0c089191f5ebdc214e0af808a\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_02878d441ceae15ce060b73daf\" ON \"drive_folder\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f4fc06e49c0171c85f1c48060d\" ON \"drive_folder\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_00ceffb0cdc238b3233294f08f\" ON \"drive_folder\" (\"parentId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"drive_file\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32), \"userHost\" character varying(128), \"md5\" character varying(32) NOT NULL, \"name\" character varying(256) NOT NULL, \"type\" character varying(128) NOT NULL, \"size\" integer NOT NULL, \"comment\" character varying(512), \"properties\" jsonb NOT NULL DEFAULT '{}', \"storedInternal\" boolean NOT NULL, \"url\" character varying(512) NOT NULL, \"thumbnailUrl\" character varying(512), \"webpublicUrl\" character varying(512), \"accessKey\" character varying(256), \"thumbnailAccessKey\" character varying(256), \"webpublicAccessKey\" character varying(256), \"uri\" character varying(512), \"src\" character varying(512), \"folderId\" character varying(32), \"isSensitive\" boolean NOT NULL DEFAULT false, \"isLink\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_43ddaaaf18c9e68029b7cbb032e\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c8dfad3b72196dd1d6b5db168a\" ON \"drive_file\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_860fa6f6c7df5bb887249fba22\" ON \"drive_file\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_92779627994ac79277f070c91e\" ON \"drive_file\" (\"userHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_37bb9a1b4585f8a3beb24c62d6\" ON \"drive_file\" (\"md5\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a40b8df8c989d7db937ea27cf6\" ON \"drive_file\" (\"type\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d85a184c2540d2deba33daf642\" ON \"drive_file\" (\"accessKey\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_e74022ce9a074b3866f70e0d27\" ON \"drive_file\" (\"thumbnailAccessKey\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_c55b2b7c284d9fef98026fc88e\" ON \"drive_file\" (\"webpublicAccessKey\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e5848eac4940934e23dbc17581\" ON \"drive_file\" (\"uri\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bb90d1956dafc4068c28aa7560\" ON \"drive_file\" (\"folderId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"updatedAt\" TIMESTAMP WITH TIME ZONE, \"lastFetchedAt\" TIMESTAMP WITH TIME ZONE, \"username\" character varying(128) NOT NULL, \"usernameLower\" character varying(128) NOT NULL, \"name\" character varying(128), \"followersCount\" integer NOT NULL DEFAULT 0, \"followingCount\" integer NOT NULL DEFAULT 0, \"notesCount\" integer NOT NULL DEFAULT 0, \"avatarId\" character varying(32), \"bannerId\" character varying(32), \"tags\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], \"avatarUrl\" character varying(512), \"bannerUrl\" character varying(512), \"avatarColor\" character varying(32), \"bannerColor\" character varying(32), \"isSuspended\" boolean NOT NULL DEFAULT false, \"isSilenced\" boolean NOT NULL DEFAULT false, \"isLocked\" boolean NOT NULL DEFAULT false, \"isBot\" boolean NOT NULL DEFAULT false, \"isCat\" boolean NOT NULL DEFAULT false, \"isAdmin\" boolean NOT NULL DEFAULT false, \"isModerator\" boolean NOT NULL DEFAULT false, \"isVerified\" boolean NOT NULL DEFAULT false, \"emojis\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], \"host\" character varying(128), \"inbox\" character varying(512), \"sharedInbox\" character varying(512), \"featured\" character varying(512), \"uri\" character varying(512), \"token\" character(16), CONSTRAINT \"UQ_a854e557b1b14814750c7c7b0c9\" UNIQUE (\"token\"), CONSTRAINT \"REL_58f5c71eaab331645112cf8cfa\" UNIQUE (\"avatarId\"), CONSTRAINT \"REL_afc64b53f8db3707ceb34eb28e\" UNIQUE (\"bannerId\"), CONSTRAINT \"PK_cace4a159ff9f2512dd42373760\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e11e649824a45d8ed01d597fd9\" ON \"user\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_80ca6e6ef65fb9ef34ea8c90f4\" ON \"user\" (\"updatedAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a27b942a0d6dcff90e3ee9b5e8\" ON \"user\" (\"usernameLower\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_fa99d777623947a5b05f394cae\" ON \"user\" (\"tags\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3252a5df8d5bbd16b281f7799e\" ON \"user\" (\"host\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_be623adaa4c566baf5d29ce0c8\" ON \"user\" (\"uri\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a854e557b1b14814750c7c7b0c\" ON \"user\" (\"token\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_5deb01ae162d1d70b80d064c27\" ON \"user\" (\"usernameLower\", \"host\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"app\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32), \"secret\" character varying(64) NOT NULL, \"name\" character varying(128) NOT NULL, \"description\" character varying(512) NOT NULL, \"permission\" character varying(64) array NOT NULL, \"callbackUrl\" character varying(512), CONSTRAINT \"PK_9478629fc093d229df09e560aea\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_048a757923ed8b157e9895da53\" ON \"app\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3f5b0899ef90527a3462d7c2cb\" ON \"app\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f49922d511d666848f250663c4\" ON \"app\" (\"secret\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"access_token\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"token\" character varying(128) NOT NULL, \"hash\" character varying(128) NOT NULL, \"userId\" character varying(32) NOT NULL, \"appId\" character varying(32) NOT NULL, CONSTRAINT \"PK_f20f028607b2603deabd8182d12\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_70ba8f6af34bc924fc9e12adb8\" ON \"access_token\" (\"token\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_64c327441248bae40f7d92f34f\" ON \"access_token\" (\"hash\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_9949557d0e1b2c19e5344c171e\" ON \"access_token\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"note_visibility_enum\" AS ENUM('public', 'home', 'followers', 'specified')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"replyId\" character varying(32), \"renoteId\" character varying(32), \"text\" text, \"name\" character varying(256), \"cw\" character varying(512), \"appId\" character varying(32), \"userId\" character varying(32) NOT NULL, \"viaMobile\" boolean NOT NULL DEFAULT false, \"localOnly\" boolean NOT NULL DEFAULT false, \"renoteCount\" smallint NOT NULL DEFAULT 0, \"repliesCount\" smallint NOT NULL DEFAULT 0, \"reactions\" jsonb NOT NULL DEFAULT '{}', \"visibility\" \"note_visibility_enum\" NOT NULL, \"uri\" character varying(512), \"score\" integer NOT NULL DEFAULT 0, \"fileIds\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[], \"attachedFileTypes\" character varying(256) array NOT NULL DEFAULT '{}'::varchar[], \"visibleUserIds\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[], \"mentions\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[], \"mentionedRemoteUsers\" text NOT NULL DEFAULT '[]', \"emojis\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], \"tags\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], \"hasPoll\" boolean NOT NULL DEFAULT false, \"geo\" jsonb DEFAULT null, \"userHost\" character varying(128), \"replyUserId\" character varying(32), \"replyUserHost\" character varying(128), \"renoteUserId\" character varying(32), \"renoteUserHost\" character varying(128), CONSTRAINT \"PK_96d0c172a4fba276b1bbed43058\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e7c0567f5261063592f022e9b5\" ON \"note\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_17cb3553c700a4985dff5a30ff\" ON \"note\" (\"replyId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_52ccc804d7c69037d558bac4c9\" ON \"note\" (\"renoteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5b87d9d19127bd5d92026017a7\" ON \"note\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_153536c67d05e9adb24e99fc2b\" ON \"note\" (\"uri\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_51c063b6a133a9cb87145450f5\" ON \"note\" (\"fileIds\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_25dfc71b0369b003a4cd434d0b\" ON \"note\" (\"attachedFileTypes\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_796a8c03959361f97dc2be1d5c\" ON \"note\" (\"visibleUserIds\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_54ebcb6d27222913b908d56fd8\" ON \"note\" (\"mentions\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_88937d94d7443d9a99a76fa5c0\" ON \"note\" (\"tags\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7125a826ab192eb27e11d358a5\" ON \"note\" (\"userHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"poll_vote\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"choice\" integer NOT NULL, CONSTRAINT \"PK_fd002d371201c472490ba89c6a0\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0fb627e1c2f753262a74f0562d\" ON \"poll_vote\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_66d2bd2ee31d14bcc23069a89f\" ON \"poll_vote\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_aecfbd5ef60374918e63ee95fa\" ON \"poll_vote\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_50bd7164c5b78f1f4a42c4d21f\" ON \"poll_vote\" (\"userId\", \"noteId\", \"choice\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note_reaction\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"reaction\" character varying(128) NOT NULL, CONSTRAINT \"PK_767ec729b108799b587a3fcc9cf\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_01f4581f114e0ebd2bbb876f0b\" ON \"note_reaction\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_13761f64257f40c5636d0ff95e\" ON \"note_reaction\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_45145e4953780f3cd5656f0ea6\" ON \"note_reaction\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_ad0c221b25672daf2df320a817\" ON \"note_reaction\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note_watching\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"noteUserId\" character varying(32) NOT NULL, CONSTRAINT \"PK_49286fdb23725945a74aa27d757\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_318cdf42a9cfc11f479bd802bb\" ON \"note_watching\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b0134ec406e8d09a540f818288\" ON \"note_watching\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_03e7028ab8388a3f5e3ce2a861\" ON \"note_watching\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_44499765eec6b5489d72c4253b\" ON \"note_watching\" (\"noteUserId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a42c93c69989ce1d09959df4cf\" ON \"note_watching\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note_unread\" (\"id\" character varying(32) NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"noteUserId\" character varying(32) NOT NULL, \"isSpecified\" boolean NOT NULL, CONSTRAINT \"PK_1904eda61a784f57e6e51fa9c1f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_56b0166d34ddae49d8ef7610bb\" ON \"note_unread\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e637cba4dc4410218c4251260e\" ON \"note_unread\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d908433a4953cc13216cd9c274\" ON \"note_unread\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"notification\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"notifieeId\" character varying(32) NOT NULL, \"notifierId\" character varying(32) NOT NULL, \"type\" character varying(32) NOT NULL, \"isRead\" boolean NOT NULL DEFAULT false, \"noteId\" character varying(32), \"reaction\" character varying(128), \"choice\" integer, CONSTRAINT \"PK_705b6c7cdf9b2c2ff7ac7872cb7\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b11a5e627c41d4dc3170f1d370\" ON \"notification\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3c601b70a1066d2c8b517094cb\" ON \"notification\" (\"notifieeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"meta\" (\"id\" character varying(32) NOT NULL, \"name\" character varying(128), \"description\" character varying(1024), \"maintainerName\" character varying(128), \"maintainerEmail\" character varying(128), \"announcements\" jsonb NOT NULL DEFAULT '[]', \"disableRegistration\" boolean NOT NULL DEFAULT false, \"disableLocalTimeline\" boolean NOT NULL DEFAULT false, \"disableGlobalTimeline\" boolean NOT NULL DEFAULT false, \"enableEmojiReaction\" boolean NOT NULL DEFAULT true, \"useStarForReactionFallback\" boolean NOT NULL DEFAULT false, \"langs\" character varying(64) array NOT NULL DEFAULT '{}'::varchar[], \"hiddenTags\" character varying(256) array NOT NULL DEFAULT '{}'::varchar[], \"blockedHosts\" character varying(256) array NOT NULL DEFAULT '{}'::varchar[], \"mascotImageUrl\" character varying(512) DEFAULT '/assets/ai.png', \"bannerUrl\" character varying(512), \"errorImageUrl\" character varying(512) DEFAULT 'https://xn--931a.moe/aiart/yubitun.png', \"iconUrl\" character varying(512), \"cacheRemoteFiles\" boolean NOT NULL DEFAULT true, \"proxyAccount\" character varying(128), \"enableRecaptcha\" boolean NOT NULL DEFAULT false, \"recaptchaSiteKey\" character varying(64), \"recaptchaSecretKey\" character varying(64), \"localDriveCapacityMb\" integer NOT NULL DEFAULT 1024, \"remoteDriveCapacityMb\" integer NOT NULL DEFAULT 32, \"maxNoteTextLength\" integer NOT NULL DEFAULT 500, \"summalyProxy\" character varying(128), \"enableEmail\" boolean NOT NULL DEFAULT false, \"email\" character varying(128), \"smtpSecure\" boolean NOT NULL DEFAULT false, \"smtpHost\" character varying(128), \"smtpPort\" integer, \"smtpUser\" character varying(128), \"smtpPass\" character varying(128), \"enableServiceWorker\" boolean NOT NULL DEFAULT false, \"swPublicKey\" character varying(128), \"swPrivateKey\" character varying(128), \"enableTwitterIntegration\" boolean NOT NULL DEFAULT false, \"twitterConsumerKey\" character varying(128), \"twitterConsumerSecret\" character varying(128), \"enableGithubIntegration\" boolean NOT NULL DEFAULT false, \"githubClientId\" character varying(128), \"githubClientSecret\" character varying(128), \"enableDiscordIntegration\" boolean NOT NULL DEFAULT false, \"discordClientId\" character varying(128), \"discordClientSecret\" character varying(128), CONSTRAINT \"PK_c4c17a6c2bd7651338b60fc590b\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"following\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"followeeId\" character varying(32) NOT NULL, \"followerId\" character varying(32) NOT NULL, \"followerHost\" character varying(128), \"followerInbox\" character varying(512), \"followerSharedInbox\" character varying(512), \"followeeHost\" character varying(128), \"followeeInbox\" character varying(512), \"followeeSharedInbox\" character varying(512), CONSTRAINT \"PK_c76c6e044bdf76ecf8bfb82a645\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_582f8fab771a9040a12961f3e7\" ON \"following\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_24e0042143a18157b234df186c\" ON \"following\" (\"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6516c5a6f3c015b4eed39978be\" ON \"following\" (\"followerId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_307be5f1d1252e0388662acb96\" ON \"following\" (\"followerId\", \"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"instance\" (\"id\" character varying(32) NOT NULL, \"caughtAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"host\" character varying(128) NOT NULL, \"system\" character varying(64), \"usersCount\" integer NOT NULL DEFAULT 0, \"notesCount\" integer NOT NULL DEFAULT 0, \"followingCount\" integer NOT NULL DEFAULT 0, \"followersCount\" integer NOT NULL DEFAULT 0, \"driveUsage\" integer NOT NULL DEFAULT 0, \"driveFiles\" integer NOT NULL DEFAULT 0, \"latestRequestSentAt\" TIMESTAMP WITH TIME ZONE, \"latestStatus\" integer, \"latestRequestReceivedAt\" TIMESTAMP WITH TIME ZONE, \"lastCommunicatedAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"isNotResponding\" boolean NOT NULL DEFAULT false, \"isMarkedAsClosed\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_eaf60e4a0c399c9935413e06474\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2cd3b2a6b4cf0b910b260afe08\" ON \"instance\" (\"caughtAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_8d5afc98982185799b160e10eb\" ON \"instance\" (\"host\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"muting\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"muteeId\" character varying(32) NOT NULL, \"muterId\" character varying(32) NOT NULL, CONSTRAINT \"PK_2e92d06c8b5c602eeb27ca9ba48\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f86d57fbca33c7a4e6897490cc\" ON \"muting\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_ec96b4fed9dae517e0dbbe0675\" ON \"muting\" (\"muteeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_93060675b4a79a577f31d260c6\" ON \"muting\" (\"muterId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_1eb9d9824a630321a29fd3b290\" ON \"muting\" (\"muterId\", \"muteeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"sw_subscription\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"endpoint\" character varying(512) NOT NULL, \"auth\" character varying(256) NOT NULL, \"publickey\" character varying(128) NOT NULL, CONSTRAINT \"PK_e8f763631530051b95eb6279b91\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_97754ca6f2baff9b4abb7f853d\" ON \"sw_subscription\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"blocking\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"blockeeId\" character varying(32) NOT NULL, \"blockerId\" character varying(32) NOT NULL, CONSTRAINT \"PK_e5d9a541cc1965ee7e048ea09dd\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b9a354f7941c1e779f3b33aea6\" ON \"blocking\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2cd4a2743a99671308f5417759\" ON \"blocking\" (\"blockeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0627125f1a8a42c9a1929edb55\" ON \"blocking\" (\"blockerId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_98a1bc5cb30dfd159de056549f\" ON \"blocking\" (\"blockerId\", \"blockeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_list\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, CONSTRAINT \"PK_87bab75775fd9b1ff822b656402\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b7fcefbdd1c18dce86687531f9\" ON \"user_list\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_list_joining\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"userListId\" character varying(32) NOT NULL, CONSTRAINT \"PK_11abb3768da1c5f8de101c9df45\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d844bfc6f3f523a05189076efa\" ON \"user_list_joining\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_605472305f26818cc93d1baaa7\" ON \"user_list_joining\" (\"userListId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"hashtag\" (\"id\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, \"mentionedUserIds\" character varying(32) array NOT NULL, \"mentionedUsersCount\" integer NOT NULL DEFAULT 0, \"mentionedLocalUserIds\" character varying(32) array NOT NULL, \"mentionedLocalUsersCount\" integer NOT NULL DEFAULT 0, \"mentionedRemoteUserIds\" character varying(32) array NOT NULL, \"mentionedRemoteUsersCount\" integer NOT NULL DEFAULT 0, \"attachedUserIds\" character varying(32) array NOT NULL, \"attachedUsersCount\" integer NOT NULL DEFAULT 0, \"attachedLocalUserIds\" character varying(32) array NOT NULL, \"attachedLocalUsersCount\" integer NOT NULL DEFAULT 0, \"attachedRemoteUserIds\" character varying(32) array NOT NULL, \"attachedRemoteUsersCount\" integer NOT NULL DEFAULT 0, CONSTRAINT \"PK_cb36eb8af8412bfa978f1165d78\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_347fec870eafea7b26c8a73bac\" ON \"hashtag\" (\"name\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2710a55f826ee236ea1a62698f\" ON \"hashtag\" (\"mentionedUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0e206cec573f1edff4a3062923\" ON \"hashtag\" (\"mentionedLocalUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4c02d38a976c3ae132228c6fce\" ON \"hashtag\" (\"mentionedRemoteUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d57f9030cd3af7f63ffb1c267c\" ON \"hashtag\" (\"attachedUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0c44bf4f680964145f2a68a341\" ON \"hashtag\" (\"attachedLocalUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0b03cbcd7e6a7ce068efa8ecc2\" ON \"hashtag\" (\"attachedRemoteUsersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note_favorite\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, CONSTRAINT \"PK_af0da35a60b9fa4463a62082b36\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_47f4b1892f5d6ba8efb3057d81\" ON \"note_favorite\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0f4fb9ad355f3effff221ef245\" ON \"note_favorite\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"abuse_user_report\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"reporterId\" character varying(32) NOT NULL, \"comment\" character varying(512) NOT NULL, CONSTRAINT \"PK_87873f5f5cc5c321a1306b2d18c\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_db2098070b2b5a523c58181f74\" ON \"abuse_user_report\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d049123c413e68ca52abe73420\" ON \"abuse_user_report\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_04cc96756f89d0b7f9473e8cdf\" ON \"abuse_user_report\" (\"reporterId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_5cd442c3b2e74fdd99dae20243\" ON \"abuse_user_report\" (\"userId\", \"reporterId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"registration_ticket\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"code\" character varying(64) NOT NULL, CONSTRAINT \"PK_f11696b6fafcf3662d4292734f8\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0ff69e8dfa9fe31bb4a4660f59\" ON \"registration_ticket\" (\"code\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"messaging_message\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"recipientId\" character varying(32) NOT NULL, \"text\" character varying(4096), \"isRead\" boolean NOT NULL DEFAULT false, \"fileId\" character varying(32), CONSTRAINT \"PK_db398fd79dc95d0eb8c30456eaa\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e21cd3646e52ef9c94aaf17c2e\" ON \"messaging_message\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5377c307783fce2b6d352e1203\" ON \"messaging_message\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_cac14a4e3944454a5ce7daa514\" ON \"messaging_message\" (\"recipientId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"signin\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"ip\" character varying(128) NOT NULL, \"headers\" jsonb NOT NULL, \"success\" boolean NOT NULL, CONSTRAINT \"PK_9e96ddc025712616fc492b3b588\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2c308dbdc50d94dc625670055f\" ON \"signin\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"auth_session\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"token\" character varying(128) NOT NULL, \"userId\" character varying(32), \"appId\" character varying(32) NOT NULL, CONSTRAINT \"PK_19354ed146424a728c1112a8cbf\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_62cb09e1129f6ec024ef66e183\" ON \"auth_session\" (\"token\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"follow_request\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"followeeId\" character varying(32) NOT NULL, \"followerId\" character varying(32) NOT NULL, \"requestId\" character varying(128), \"followerHost\" character varying(128), \"followerInbox\" character varying(512), \"followerSharedInbox\" character varying(512), \"followeeHost\" character varying(128), \"followeeInbox\" character varying(512), \"followeeSharedInbox\" character varying(512), CONSTRAINT \"PK_53a9aa3725f7a3deb150b39dbfc\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_12c01c0d1a79f77d9f6c15fadd\" ON \"follow_request\" (\"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a7fd92dd6dc519e6fb435dd108\" ON \"follow_request\" (\"followerId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d54a512b822fac7ed52800f6b4\" ON \"follow_request\" (\"followerId\", \"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"emoji\" (\"id\" character varying(32) NOT NULL, \"updatedAt\" TIMESTAMP WITH TIME ZONE, \"name\" character varying(128) NOT NULL, \"host\" character varying(128), \"url\" character varying(512) NOT NULL, \"uri\" character varying(512), \"type\" character varying(64), \"aliases\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], CONSTRAINT \"PK_df74ce05e24999ee01ea0bc50a3\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b37dafc86e9af007e3295c2781\" ON \"emoji\" (\"name\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5900e907bb46516ddf2871327c\" ON \"emoji\" (\"host\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_4f4d35e1256c84ae3d1f0eab10\" ON \"emoji\" (\"name\", \"host\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"reversi_game\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"startedAt\" TIMESTAMP WITH TIME ZONE, \"user1Id\" character varying(32) NOT NULL, \"user2Id\" character varying(32) NOT NULL, \"user1Accepted\" boolean NOT NULL DEFAULT false, \"user2Accepted\" boolean NOT NULL DEFAULT false, \"black\" integer, \"isStarted\" boolean NOT NULL DEFAULT false, \"isEnded\" boolean NOT NULL DEFAULT false, \"winnerId\" character varying(32), \"surrendered\" character varying(32), \"logs\" jsonb NOT NULL DEFAULT '[]', \"map\" character varying(64) array NOT NULL, \"bw\" character varying(32) NOT NULL, \"isLlotheo\" boolean NOT NULL DEFAULT false, \"canPutEverywhere\" boolean NOT NULL DEFAULT false, \"loopedBoard\" boolean NOT NULL DEFAULT false, \"form1\" jsonb DEFAULT null, \"form2\" jsonb DEFAULT null, \"crc32\" character varying(32), CONSTRAINT \"PK_76b30eeba71b1193ad7c5311c3f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b46ec40746efceac604142be1c\" ON \"reversi_game\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"reversi_matching\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"parentId\" character varying(32) NOT NULL, \"childId\" character varying(32) NOT NULL, CONSTRAINT \"PK_880bd0afbab232f21c8b9d146cf\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b604d92d6c7aec38627f6eaf16\" ON \"reversi_matching\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3b25402709dd9882048c2bbade\" ON \"reversi_matching\" (\"parentId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e247b23a3c9b45f89ec1299d06\" ON \"reversi_matching\" (\"childId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_note_pining\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, CONSTRAINT \"PK_a6a2dad4ae000abce2ea9d9b103\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bfbc6f79ba4007b4ce5097f08d\" ON \"user_note_pining\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_410cd649884b501c02d6e72738\" ON \"user_note_pining\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"poll_notevisibility_enum\" AS ENUM('public', 'home', 'followers', 'specified')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"poll\" (\"noteId\" character varying(32) NOT NULL, \"expiresAt\" TIMESTAMP WITH TIME ZONE, \"multiple\" boolean NOT NULL, \"choices\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], \"votes\" integer array NOT NULL, \"noteVisibility\" \"poll_notevisibility_enum\" NOT NULL, \"userId\" character varying(32) NOT NULL, \"userHost\" character varying(128), CONSTRAINT \"REL_da851e06d0dfe2ef397d8b1bf1\" UNIQUE (\"noteId\"), CONSTRAINT \"PK_da851e06d0dfe2ef397d8b1bf1b\" PRIMARY KEY (\"noteId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0610ebcfcfb4a18441a9bcdab2\" ON \"poll\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7fa20a12319c7f6dc3aed98c0a\" ON \"poll\" (\"userHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_keypair\" (\"userId\" character varying(32) NOT NULL, \"publicKey\" character varying(4096) NOT NULL, \"privateKey\" character varying(4096) NOT NULL, CONSTRAINT \"REL_f4853eb41ab722fe05f81cedeb\" UNIQUE (\"userId\"), CONSTRAINT \"PK_f4853eb41ab722fe05f81cedeb6\" PRIMARY KEY (\"userId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_publickey\" (\"userId\" character varying(32) NOT NULL, \"keyId\" character varying(256) NOT NULL, \"keyPem\" character varying(4096) NOT NULL, CONSTRAINT \"REL_10c146e4b39b443ede016f6736\" UNIQUE (\"userId\"), CONSTRAINT \"PK_10c146e4b39b443ede016f6736d\" PRIMARY KEY (\"userId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_171e64971c780ebd23fae140bb\" ON \"user_publickey\" (\"keyId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_profile\" (\"userId\" character varying(32) NOT NULL, \"location\" character varying(128), \"birthday\" character(10), \"description\" character varying(1024), \"fields\" jsonb NOT NULL DEFAULT '[]', \"url\" character varying(512), \"email\" character varying(128), \"emailVerifyCode\" character varying(128), \"emailVerified\" boolean NOT NULL DEFAULT false, \"twoFactorTempSecret\" character varying(128), \"twoFactorSecret\" character varying(128), \"twoFactorEnabled\" boolean NOT NULL DEFAULT false, \"password\" character varying(128), \"clientData\" jsonb NOT NULL DEFAULT '{}', \"autoWatch\" boolean NOT NULL DEFAULT false, \"autoAcceptFollowed\" boolean NOT NULL DEFAULT false, \"alwaysMarkNsfw\" boolean NOT NULL DEFAULT false, \"carefulBot\" boolean NOT NULL DEFAULT false, \"twitter\" boolean NOT NULL DEFAULT false, \"twitterAccessToken\" character varying(64) DEFAULT null, \"twitterAccessTokenSecret\" character varying(64) DEFAULT null, \"twitterUserId\" character varying(64) DEFAULT null, \"twitterScreenName\" character varying(64) DEFAULT null, \"github\" boolean NOT NULL DEFAULT false, \"githubAccessToken\" character varying(64) DEFAULT null, \"githubId\" integer DEFAULT null, \"githubLogin\" character varying(64) DEFAULT null, \"discord\" boolean NOT NULL DEFAULT false, \"discordAccessToken\" character varying(64) DEFAULT null, \"discordRefreshToken\" character varying(64) DEFAULT null, \"discordExpiresDate\" integer DEFAULT null, \"discordId\" character varying(64) DEFAULT null, \"discordUsername\" character varying(64) DEFAULT null, \"discordDiscriminator\" character varying(64) DEFAULT null, \"userHost\" character varying(128), CONSTRAINT \"REL_51cb79b5555effaf7d69ba1cff\" UNIQUE (\"userId\"), CONSTRAINT \"PK_51cb79b5555effaf7d69ba1cff9\" PRIMARY KEY (\"userId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_dce530b98e454793dac5ec2f5a\" ON \"user_profile\" (\"userHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__active_users_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__active_users\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__active_users_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_count\" bigint NOT NULL, \"___remote_count\" bigint NOT NULL, CONSTRAINT \"PK_317237a9f733b970604a11e314f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__drive_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__drive\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__drive_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_totalCount\" bigint NOT NULL, \"___local_totalSize\" bigint NOT NULL, \"___local_incCount\" bigint NOT NULL, \"___local_incSize\" bigint NOT NULL, \"___local_decCount\" bigint NOT NULL, \"___local_decSize\" bigint NOT NULL, \"___remote_totalCount\" bigint NOT NULL, \"___remote_totalSize\" bigint NOT NULL, \"___remote_incCount\" bigint NOT NULL, \"___remote_incSize\" bigint NOT NULL, \"___remote_decCount\" bigint NOT NULL, \"___remote_decSize\" bigint NOT NULL, CONSTRAINT \"PK_f96bc548a765cd4b3b354221ce7\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__federation_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__federation\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__federation_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___instance_total\" bigint NOT NULL, \"___instance_inc\" bigint NOT NULL, \"___instance_dec\" bigint NOT NULL, CONSTRAINT \"PK_b39dcd31a0fe1a7757e348e85fd\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__hashtag_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__hashtag\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__hashtag_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_count\" bigint NOT NULL, \"___remote_count\" bigint NOT NULL, CONSTRAINT \"PK_c32f1ea2b44a5d2f7881e37f8f9\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__instance_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__instance\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__instance_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___requests_failed\" bigint NOT NULL, \"___requests_succeeded\" bigint NOT NULL, \"___requests_received\" bigint NOT NULL, \"___notes_total\" bigint NOT NULL, \"___notes_inc\" bigint NOT NULL, \"___notes_dec\" bigint NOT NULL, \"___notes_diffs_normal\" bigint NOT NULL, \"___notes_diffs_reply\" bigint NOT NULL, \"___notes_diffs_renote\" bigint NOT NULL, \"___users_total\" bigint NOT NULL, \"___users_inc\" bigint NOT NULL, \"___users_dec\" bigint NOT NULL, \"___following_total\" bigint NOT NULL, \"___following_inc\" bigint NOT NULL, \"___following_dec\" bigint NOT NULL, \"___followers_total\" bigint NOT NULL, \"___followers_inc\" bigint NOT NULL, \"___followers_dec\" bigint NOT NULL, \"___drive_totalFiles\" bigint NOT NULL, \"___drive_totalUsage\" bigint NOT NULL, \"___drive_incFiles\" bigint NOT NULL, \"___drive_incUsage\" bigint NOT NULL, \"___drive_decFiles\" bigint NOT NULL, \"___drive_decUsage\" bigint NOT NULL, CONSTRAINT \"PK_1267c67c7c2d47b4903975f2c00\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__network_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__network\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__network_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___incomingRequests\" bigint NOT NULL, \"___outgoingRequests\" bigint NOT NULL, \"___totalTime\" bigint NOT NULL, \"___incomingBytes\" bigint NOT NULL, \"___outgoingBytes\" bigint NOT NULL, CONSTRAINT \"PK_bc4290c2e27fad14ef0c1ca93f3\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__notes_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__notes\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__notes_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_total\" bigint NOT NULL, \"___local_inc\" bigint NOT NULL, \"___local_dec\" bigint NOT NULL, \"___local_diffs_normal\" bigint NOT NULL, \"___local_diffs_reply\" bigint NOT NULL, \"___local_diffs_renote\" bigint NOT NULL, \"___remote_total\" bigint NOT NULL, \"___remote_inc\" bigint NOT NULL, \"___remote_dec\" bigint NOT NULL, \"___remote_diffs_normal\" bigint NOT NULL, \"___remote_diffs_reply\" bigint NOT NULL, \"___remote_diffs_renote\" bigint NOT NULL, CONSTRAINT \"PK_0aec823fa85c7f901bdb3863b14\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__per_user_drive_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__per_user_drive\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__per_user_drive_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___totalCount\" bigint NOT NULL, \"___totalSize\" bigint NOT NULL, \"___incCount\" bigint NOT NULL, \"___incSize\" bigint NOT NULL, \"___decCount\" bigint NOT NULL, \"___decSize\" bigint NOT NULL, CONSTRAINT \"PK_d0ef23d24d666e1a44a0cd3d208\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__per_user_following_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__per_user_following\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__per_user_following_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_followings_total\" bigint NOT NULL, \"___local_followings_inc\" bigint NOT NULL, \"___local_followings_dec\" bigint NOT NULL, \"___local_followers_total\" bigint NOT NULL, \"___local_followers_inc\" bigint NOT NULL, \"___local_followers_dec\" bigint NOT NULL, \"___remote_followings_total\" bigint NOT NULL, \"___remote_followings_inc\" bigint NOT NULL, \"___remote_followings_dec\" bigint NOT NULL, \"___remote_followers_total\" bigint NOT NULL, \"___remote_followers_inc\" bigint NOT NULL, \"___remote_followers_dec\" bigint NOT NULL, CONSTRAINT \"PK_85bb1b540363a29c2fec83bd907\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__per_user_notes_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__per_user_notes\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__per_user_notes_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___total\" bigint NOT NULL, \"___inc\" bigint NOT NULL, \"___dec\" bigint NOT NULL, \"___diffs_normal\" bigint NOT NULL, \"___diffs_reply\" bigint NOT NULL, \"___diffs_renote\" bigint NOT NULL, CONSTRAINT \"PK_334acf6e915af2f29edc11b8e50\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__per_user_reaction_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__per_user_reaction\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__per_user_reaction_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_count\" bigint NOT NULL, \"___remote_count\" bigint NOT NULL, CONSTRAINT \"PK_984f54dae441e65b633e8d27a7f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__test_grouped_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__test_grouped\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__test_grouped_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___foo_total\" bigint NOT NULL, \"___foo_inc\" bigint NOT NULL, \"___foo_dec\" bigint NOT NULL, CONSTRAINT \"PK_f4a2b175d308695af30d4293272\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__test_unique_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__test_unique\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__test_unique_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___foo\" bigint NOT NULL, CONSTRAINT \"PK_409bac9c97cc612d8500012319d\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__test_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__test\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__test_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___foo_total\" bigint NOT NULL, \"___foo_inc\" bigint NOT NULL, \"___foo_dec\" bigint NOT NULL, CONSTRAINT \"PK_b4bc31dffbd1b785276a3ecfc1e\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"__chart__users_span_enum\" AS ENUM('hour', 'day')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__users\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128), \"span\" \"__chart__users_span_enum\" NOT NULL, \"unique\" jsonb NOT NULL DEFAULT '{}', \"___local_total\" bigint NOT NULL, \"___local_inc\" bigint NOT NULL, \"___local_dec\" bigint NOT NULL, \"___remote_total\" bigint NOT NULL, \"___remote_inc\" bigint NOT NULL, \"___remote_dec\" bigint NOT NULL, CONSTRAINT \"PK_4dfcf2c78d03524b9eb2c99d328\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_folder\" ADD CONSTRAINT \"FK_f4fc06e49c0171c85f1c48060d2\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_folder\" ADD CONSTRAINT \"FK_00ceffb0cdc238b3233294f08f2\" FOREIGN KEY (\"parentId\") REFERENCES \"drive_folder\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD CONSTRAINT \"FK_860fa6f6c7df5bb887249fba22e\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD CONSTRAINT \"FK_bb90d1956dafc4068c28aa7560a\" FOREIGN KEY (\"folderId\") REFERENCES \"drive_folder\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD CONSTRAINT \"FK_58f5c71eaab331645112cf8cfa5\" FOREIGN KEY (\"avatarId\") REFERENCES \"drive_file\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD CONSTRAINT \"FK_afc64b53f8db3707ceb34eb28e2\" FOREIGN KEY (\"bannerId\") REFERENCES \"drive_file\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"app\" ADD CONSTRAINT \"FK_3f5b0899ef90527a3462d7c2cb3\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD CONSTRAINT \"FK_9949557d0e1b2c19e5344c171e9\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD CONSTRAINT \"FK_a3ff16c90cc87a82a0b5959e560\" FOREIGN KEY (\"appId\") REFERENCES \"app\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD CONSTRAINT \"FK_17cb3553c700a4985dff5a30ff5\" FOREIGN KEY (\"replyId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD CONSTRAINT \"FK_52ccc804d7c69037d558bac4c96\" FOREIGN KEY (\"renoteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD CONSTRAINT \"FK_ec5c201576192ba8904c345c5cc\" FOREIGN KEY (\"appId\") REFERENCES \"app\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD CONSTRAINT \"FK_5b87d9d19127bd5d92026017a7b\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"poll_vote\" ADD CONSTRAINT \"FK_66d2bd2ee31d14bcc23069a89f8\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"poll_vote\" ADD CONSTRAINT \"FK_aecfbd5ef60374918e63ee95fa7\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_reaction\" ADD CONSTRAINT \"FK_13761f64257f40c5636d0ff95ee\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_reaction\" ADD CONSTRAINT \"FK_45145e4953780f3cd5656f0ea6a\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_watching\" ADD CONSTRAINT \"FK_b0134ec406e8d09a540f8182888\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_watching\" ADD CONSTRAINT \"FK_03e7028ab8388a3f5e3ce2a8619\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" ADD CONSTRAINT \"FK_56b0166d34ddae49d8ef7610bb9\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" ADD CONSTRAINT \"FK_e637cba4dc4410218c4251260e4\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_3c601b70a1066d2c8b517094cb9\" FOREIGN KEY (\"notifieeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_3b4e96eec8d36a8bbb9d02aa710\" FOREIGN KEY (\"notifierId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_769cb6b73a1efe22ddf733ac453\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"following\" ADD CONSTRAINT \"FK_24e0042143a18157b234df186c3\" FOREIGN KEY (\"followeeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"following\" ADD CONSTRAINT \"FK_6516c5a6f3c015b4eed39978be5\" FOREIGN KEY (\"followerId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muting\" ADD CONSTRAINT \"FK_ec96b4fed9dae517e0dbbe0675c\" FOREIGN KEY (\"muteeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muting\" ADD CONSTRAINT \"FK_93060675b4a79a577f31d260c67\" FOREIGN KEY (\"muterId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"sw_subscription\" ADD CONSTRAINT \"FK_97754ca6f2baff9b4abb7f853dd\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"blocking\" ADD CONSTRAINT \"FK_2cd4a2743a99671308f5417759e\" FOREIGN KEY (\"blockeeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"blocking\" ADD CONSTRAINT \"FK_0627125f1a8a42c9a1929edb552\" FOREIGN KEY (\"blockerId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_list\" ADD CONSTRAINT \"FK_b7fcefbdd1c18dce86687531f99\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_list_joining\" ADD CONSTRAINT \"FK_d844bfc6f3f523a05189076efaa\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_list_joining\" ADD CONSTRAINT \"FK_605472305f26818cc93d1baaa74\" FOREIGN KEY (\"userListId\") REFERENCES \"user_list\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_favorite\" ADD CONSTRAINT \"FK_47f4b1892f5d6ba8efb3057d81a\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_favorite\" ADD CONSTRAINT \"FK_0e00498f180193423c992bc4370\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD CONSTRAINT \"FK_d049123c413e68ca52abe734203\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD CONSTRAINT \"FK_04cc96756f89d0b7f9473e8cdf3\" FOREIGN KEY (\"reporterId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD CONSTRAINT \"FK_5377c307783fce2b6d352e1203b\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD CONSTRAINT \"FK_cac14a4e3944454a5ce7daa5142\" FOREIGN KEY (\"recipientId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD CONSTRAINT \"FK_535def119223ac05ad3fa9ef64b\" FOREIGN KEY (\"fileId\") REFERENCES \"drive_file\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"signin\" ADD CONSTRAINT \"FK_2c308dbdc50d94dc625670055f7\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"auth_session\" ADD CONSTRAINT \"FK_c072b729d71697f959bde66ade0\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"auth_session\" ADD CONSTRAINT \"FK_dbe037d4bddd17b03a1dc778dee\" FOREIGN KEY (\"appId\") REFERENCES \"app\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"follow_request\" ADD CONSTRAINT \"FK_12c01c0d1a79f77d9f6c15fadd2\" FOREIGN KEY (\"followeeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"follow_request\" ADD CONSTRAINT \"FK_a7fd92dd6dc519e6fb435dd108f\" FOREIGN KEY (\"followerId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_game\" ADD CONSTRAINT \"FK_f7467510c60a45ce5aca6292743\" FOREIGN KEY (\"user1Id\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_game\" ADD CONSTRAINT \"FK_6649a4e8c5d5cf32fb03b5da9f6\" FOREIGN KEY (\"user2Id\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_matching\" ADD CONSTRAINT \"FK_3b25402709dd9882048c2bbade0\" FOREIGN KEY (\"parentId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"reversi_matching\" ADD CONSTRAINT \"FK_e247b23a3c9b45f89ec1299d066\" FOREIGN KEY (\"childId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_note_pining\" ADD CONSTRAINT \"FK_bfbc6f79ba4007b4ce5097f08d6\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_note_pining\" ADD CONSTRAINT \"FK_68881008f7c3588ad7ecae471cf\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"poll\" ADD CONSTRAINT \"FK_da851e06d0dfe2ef397d8b1bf1b\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_keypair\" ADD CONSTRAINT \"FK_f4853eb41ab722fe05f81cedeb6\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_publickey\" ADD CONSTRAINT \"FK_10c146e4b39b443ede016f6736d\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD CONSTRAINT \"FK_51cb79b5555effaf7d69ba1cff9\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"page_visibility_enum\" AS ENUM('public', 'followers', 'specified')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"page\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"updatedAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"title\" character varying(256) NOT NULL, \"name\" character varying(256) NOT NULL, \"summary\" character varying(256), \"alignCenter\" boolean NOT NULL, \"font\" character varying(32) NOT NULL, \"userId\" character varying(32) NOT NULL, \"eyeCatchingImageId\" character varying(32), \"content\" jsonb NOT NULL DEFAULT '[]', \"variables\" jsonb NOT NULL DEFAULT '[]', \"visibility\" \"page_visibility_enum\" NOT NULL, \"visibleUserIds\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[], CONSTRAINT \"PK_742f4117e065c5b6ad21b37ba1f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_fbb4297c927a9b85e9cefa2eb1\" ON \"page\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_af639b066dfbca78b01a920f8a\" ON \"page\" (\"updatedAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b82c19c08afb292de4600d99e4\" ON \"page\" (\"name\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_ae1d917992dd0c9d9bbdad06c4\" ON \"page\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_90148bbc2bf0854428786bfc15\" ON \"page\" (\"visibleUserIds\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_2133ef8317e4bdb839c0dcbf13\" ON \"page\" (\"userId\", \"name\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" ADD CONSTRAINT \"FK_ae1d917992dd0c9d9bbdad06c4a\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" ADD CONSTRAINT \"FK_3126dd7c502c9e4d7597ef7ef10\" FOREIGN KEY (\"eyeCatchingImageId\") REFERENCES \"drive_file\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ALTER COLUMN \"githubId\" TYPE VARCHAR(64) USING \"githubId\"::VARCHAR(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ALTER COLUMN \"discordExpiresDate\" TYPE VARCHAR(64) USING \"discordExpiresDate\"::VARCHAR(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"pinnedUsers\" character varying(256) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"ToSUrl\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"repositoryUrl\" character varying(512) NOT NULL DEFAULT 'https://github.com/misskey-dev/misskey'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"feedbackUrl\" character varying(512) DEFAULT 'https://github.com/misskey-dev/misskey/issues/new'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"useObjectStorage\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageBucket\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStoragePrefix\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageBaseUrl\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageEndpoint\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageRegion\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageAccessKey\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageSecretKey\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStoragePort\" integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageUseSSL\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"page_like\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"pageId\" character varying(32) NOT NULL, CONSTRAINT \"PK_813f034843af992d3ae0f43c64c\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0e61efab7f88dbb79c9166dbb4\" ON \"page_like\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_4ce6fb9c70529b4c8ac46c9bfa\" ON \"page_like\" (\"userId\", \"pageId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" ADD \"likedCount\" integer NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page_like\" ADD CONSTRAINT \"FK_0e61efab7f88dbb79c9166dbb48\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page_like\" ADD CONSTRAINT \"FK_cf8782626dced3176038176a847\" FOREIGN KEY (\"pageId\") REFERENCES \"page\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_group\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"name\" character varying(256) NOT NULL, \"userId\" character varying(32) NOT NULL, \"isPrivate\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_3c29fba6fe013ec8724378ce7c9\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_20e30aa35180e317e133d75316\" ON \"user_group\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3d6b372788ab01be58853003c9\" ON \"user_group\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_group_joining\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"userGroupId\" character varying(32) NOT NULL, CONSTRAINT \"PK_15f2425885253c5507e1599cfe7\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f3a1b4bd0c7cabba958a0c0b23\" ON \"user_group_joining\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_67dc758bc0566985d1b3d39986\" ON \"user_group_joining\" (\"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD \"groupId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD \"reads\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ALTER COLUMN \"recipientId\" DROP NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"messaging_message\".\"recipientId\" IS 'The recipient user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2c4be03b446884f9e9c502135b\" ON \"messaging_message\" (\"groupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD CONSTRAINT \"FK_2c4be03b446884f9e9c502135be\" FOREIGN KEY (\"groupId\") REFERENCES \"user_group\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group\" ADD CONSTRAINT \"FK_3d6b372788ab01be58853003c93\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_joining\" ADD CONSTRAINT \"FK_f3a1b4bd0c7cabba958a0c0b231\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_joining\" ADD CONSTRAINT \"FK_67dc758bc0566985d1b3d399865\" FOREIGN KEY (\"userGroupId\") REFERENCES \"user_group\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_group_invite\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"userGroupId\" character varying(32) NOT NULL, CONSTRAINT \"PK_3893884af0d3a5f4d01e7921a97\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_1039988afa3bf991185b277fe0\" ON \"user_group_invite\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e10924607d058004304611a436\" ON \"user_group_invite\" (\"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_78787741f9010886796f2320a4\" ON \"user_group_invite\" (\"userId\", \"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d9ecaed8c6dc43f3592c229282\" ON \"user_group_joining\" (\"userId\", \"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_invite\" ADD CONSTRAINT \"FK_1039988afa3bf991185b277fe03\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_invite\" ADD CONSTRAINT \"FK_e10924607d058004304611a436a\" FOREIGN KEY (\"userGroupId\") REFERENCES \"user_group\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_90f7da835e4c10aca6853621e1\" ON \"user_list_joining\" (\"userId\", \"userListId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"attestation_challenge\" (\"id\" character varying(32) NOT NULL, \"userId\" character varying(32) NOT NULL, \"challenge\" character varying(64) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"registrationChallenge\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_d0ba6786e093f1bcb497572a6b5\" PRIMARY KEY (\"id\", \"userId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f1a461a618fa1755692d0e0d59\" ON \"attestation_challenge\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_47efb914aed1f72dd39a306c7b\" ON \"attestation_challenge\" (\"challenge\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_security_key\" (\"id\" character varying NOT NULL, \"userId\" character varying(32) NOT NULL, \"publicKey\" character varying NOT NULL, \"lastUsed\" TIMESTAMP WITH TIME ZONE NOT NULL, \"name\" character varying(30) NOT NULL, CONSTRAINT \"PK_3e508571121ab39c5f85d10c166\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_ff9ca3b5f3ee3d0681367a9b44\" ON \"user_security_key\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0d7718e562dcedd0aa5cf2c9f7\" ON \"user_security_key\" (\"publicKey\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"securityKeysAvailable\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"attestation_challenge\" ADD CONSTRAINT \"FK_f1a461a618fa1755692d0e0d592\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_security_key\" ADD CONSTRAINT \"FK_ff9ca3b5f3ee3d0681367a9b447\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0ad37b7ef50f4ddc84363d7ccc\" ON \"__chart__active_users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_15e91a03aeeac9dbccdf43fc06\" ON \"__chart__active_users\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_00ed5f86db1f7efafb1978bf21\" ON \"__chart__active_users\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_20f57cc8f142c131340ee16742\" ON \"__chart__active_users\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_9a3ed15a30ab7e3a37702e6e08\" ON \"__chart__active_users\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c26e2c1cbb6e911e0554b27416\" ON \"__chart__active_users\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_13565815f618a1ff53886c5b28\" ON \"__chart__drive\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3fa0d0f17ca72e3dc80999a032\" ON \"__chart__drive\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7a170f67425e62a8fabb76c872\" ON \"__chart__drive\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6e1df243476e20cbf86572ecc0\" ON \"__chart__drive\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3313d7288855ec105b5bbf6c21\" ON \"__chart__drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_06690fc959f1c9fdaf21928222\" ON \"__chart__drive\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_36cb699c49580d4e6c2e6159f9\" ON \"__chart__federation\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e447064455928cf627590ef527\" ON \"__chart__federation\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_76e87c7bfc5d925fcbba405d84\" ON \"__chart__federation\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2d416e6af791a82e338c79d480\" ON \"__chart__federation\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_dd907becf76104e4b656659e6b\" ON \"__chart__federation\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e9cd07672b37d8966cf3709283\" ON \"__chart__federation\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_07747a1038c05f532a718fe1de\" ON \"__chart__hashtag\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_fcc181fb8283009c61cc4083ef\" ON \"__chart__hashtag\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_99a7d2faaef84a6f728d714ad6\" ON \"__chart__hashtag\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_49975586f50ed7b800fdd88fbd\" ON \"__chart__hashtag\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_25a97c02003338124b2b75fdbc\" ON \"__chart__hashtag\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6d6f156ceefc6bc5f273a0e370\" ON \"__chart__hashtag\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6b8f34a1a64b06014b6fb66824\" ON \"__chart__instance\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c12f0af4a66cdd30c2287ce8aa\" ON \"__chart__instance\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_da8a46ba84ca1d8bb5a29bfb63\" ON \"__chart__instance\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d0a4f79af5a97b08f37b547197\" ON \"__chart__instance\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_39ee857ab2f23493037c6b6631\" ON \"__chart__instance\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f5448d9633cff74208d850aabe\" ON \"__chart__instance\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a1efd3e0048a5f2793a47360dc\" ON \"__chart__network\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f8dd01baeded2ffa833e0a610a\" ON \"__chart__network\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7b5da130992ec9df96712d4290\" ON \"__chart__network\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_08fac0eb3b11f04c200c0b40dd\" ON \"__chart__network\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0a905b992fecd2b5c3fb98759e\" ON \"__chart__network\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_9ff6944f01acb756fdc92d7563\" ON \"__chart__network\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_42eb716a37d381cdf566192b2b\" ON \"__chart__notes\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e69096589f11e3baa98ddd64d0\" ON \"__chart__notes\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7036f2957151588b813185c794\" ON \"__chart__notes\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0c9a159c5082cbeef3ca6706b5\" ON \"__chart__notes\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f09d543e3acb16c5976bdb31fa\" ON \"__chart__notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_924fc196c80ca24bae01dd37e4\" ON \"__chart__notes\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5f86db6492274e07c1a3cdf286\" ON \"__chart__per_user_drive\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_328f259961e60c4fa0bfcf55ca\" ON \"__chart__per_user_drive\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e496ca8096d28f6b9b509264dc\" ON \"__chart__per_user_drive\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_42ea9381f0fda8dfe0fa1c8b53\" ON \"__chart__per_user_drive\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_30bf67687f483ace115c5ca642\" ON \"__chart__per_user_drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f2aeafde2ae6fbad38e857631b\" ON \"__chart__per_user_drive\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7af07790712aa3438ff6773f3b\" ON \"__chart__per_user_following\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f92dd6d03f8d994f29987f6214\" ON \"__chart__per_user_following\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4b3593098b6edc9c5afe36b18b\" ON \"__chart__per_user_following\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_57b5458d0d3d6d1e7f13d4e57f\" ON \"__chart__per_user_following\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b77d4dd9562c3a899d9a286fcd\" ON \"__chart__per_user_following\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4db3b84c7be0d3464714f3e0b1\" ON \"__chart__per_user_following\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_84234bd1abb873f07329681c83\" ON \"__chart__per_user_notes\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8d2cbbc8114d90d19b44d626b6\" ON \"__chart__per_user_notes\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_55bf20f366979f2436de99206b\" ON \"__chart__per_user_notes\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_046feeb12e9ef5f783f409866a\" ON \"__chart__per_user_notes\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5048e9daccbbbc6d567bb142d3\" ON \"__chart__per_user_notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f68a5ab958f9f5fa17a32ac23b\" ON \"__chart__per_user_notes\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f7bf4c62059764c2c2bb40fdab\" ON \"__chart__per_user_reaction\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_65633a106bce43fc7c5c30a5c7\" ON \"__chart__per_user_reaction\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8cf3156fd7a6b15c43459c6e3b\" ON \"__chart__per_user_reaction\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_edeb73c09c3143a81bcb34d569\" ON \"__chart__per_user_reaction\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_229a41ad465f9205f1f5703291\" ON \"__chart__per_user_reaction\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e316f01a6d24eb31db27f88262\" ON \"__chart__per_user_reaction\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0c641990ecf47d2545df4edb75\" ON \"__chart__test_grouped\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2be7ec6cebddc14dc11e206686\" ON \"__chart__test_grouped\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_234dff3c0b56a6150b95431ab9\" ON \"__chart__test_grouped\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a5133470f4825902e170328ca5\" ON \"__chart__test_grouped\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b14489029e4b3aaf4bba5fb524\" ON \"__chart__test_grouped\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_84e661abb7bd1e51b690d4b017\" ON \"__chart__test_grouped\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_437bab3c6061d90f6bb65fd2cc\" ON \"__chart__test_unique\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5c73bf61da4f6e6f15bae88ed1\" ON \"__chart__test_unique\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bbfa573a8181018851ed0b6357\" ON \"__chart__test_unique\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d70c86baedc68326be11f9c0ce\" ON \"__chart__test_unique\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a0cd75442dd10d0643a17c4a49\" ON \"__chart__test_unique\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_66e1e1ecd2f29e57778af35b59\" ON \"__chart__test_unique\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_b070a906db04b44c67c6c2144d\" ON \"__chart__test\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_92255988735563f0fe4aba1f05\" ON \"__chart__test\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d41cce6aee1a50bfc062038f9b\" ON \"__chart__test\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c5870993e25c3d5771f91f5003\" ON \"__chart__test\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a319e5dbf47e8a17497623beae\" ON \"__chart__test\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f170de677ea75ad4533de2723e\" ON \"__chart__test\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_845254b3eaf708ae8a6cac3026\" ON \"__chart__users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7c184198ecf66a8d3ecb253ab3\" ON \"__chart__users\" (\"span\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_ed9b95919c672a13008e9487ee\" ON \"__chart__users\" (\"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f091abb24193d50c653c6b77fc\" ON \"__chart__users\" (\"span\", \"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_337e9599f278bd7537fe30876f\" ON \"__chart__users\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a770a57c70e668cc61590c9161\" ON \"__chart__users\" (\"span\", \"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD COLUMN \"usePasswordLessLogin\" boolean DEFAULT false NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"pinnedPageId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD CONSTRAINT \"UQ_6dc44f1ceb65b1e72bacef2ca27\" UNIQUE (\"pinnedPageId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD CONSTRAINT \"FK_6dc44f1ceb65b1e72bacef2ca27\" FOREIGN KEY (\"pinnedPageId\") REFERENCES \"page\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" ADD \"hideTitleWhenPinned\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"moderation_log\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"type\" character varying(128) NOT NULL, \"info\" jsonb NOT NULL, CONSTRAINT \"PK_d0adca6ecfd068db83e4526cc26\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a08ad074601d204e0f69da9a95\" ON \"moderation_log\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"moderation_log\" ADD CONSTRAINT \"FK_a08ad074601d204e0f69da9a954\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"used_username\" (\"username\" character varying(128) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, CONSTRAINT \"PK_78fd79d2d24c6ac2f4cc9a31a5d\" PRIMARY KEY (\"username\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"room\" jsonb NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"emoji\" ADD \"category\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" DROP COLUMN \"system\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"softwareName\" character varying(64) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"softwareVersion\" character varying(64) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"openRegistrations\" boolean DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"name\" character varying(256) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"description\" character varying(4096) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"maintainerName\" character varying(128) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"maintainerEmail\" character varying(256) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"infoUpdatedAt\" TIMESTAMP WITH TIME ZONE"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"messaging_message\" ADD \"uri\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"proxyRemoteFiles\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"announcement\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"text\" character varying(8192) NOT NULL, \"title\" character varying(256) NOT NULL, \"imageUrl\" character varying(1024), CONSTRAINT \"PK_e0ef0550174fd1099a308fd18a0\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_118ec703e596086fc4515acb39\" ON \"announcement\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"announcement_read\" (\"id\" character varying(32) NOT NULL, \"userId\" character varying(32) NOT NULL, \"announcementId\" character varying(32) NOT NULL, CONSTRAINT \"PK_4b90ad1f42681d97b2683890c5e\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8288151386172b8109f7239ab2\" ON \"announcement_read\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_603a7b1e7aa0533c6c88e9bfaf\" ON \"announcement_read\" (\"announcementId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_924fa71815cfa3941d003702a0\" ON \"announcement_read\" (\"userId\", \"announcementId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"isVerified\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" DROP COLUMN \"announcements\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" DROP COLUMN \"enableEmojiReaction\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"announcement_read\" ADD CONSTRAINT \"FK_8288151386172b8109f7239ab28\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"announcement_read\" ADD CONSTRAINT \"FK_603a7b1e7aa0533c6c88e9bfafe\" FOREIGN KEY (\"announcementId\") REFERENCES \"announcement\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"announcement_read\" ADD \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"announcement\" ADD \"updatedAt\" TIMESTAMP WITH TIME ZONE"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"followRequestId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_bd7fab507621e635b32cd31892c\" FOREIGN KEY (\"followRequestId\") REFERENCES \"follow_request\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"clip\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, \"isPublic\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_f0685dac8d4dd056d7255670b75\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2b5ec6c574d6802c94c80313fb\" ON \"clip\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"clip_note\" (\"id\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"clipId\" character varying(32) NOT NULL, CONSTRAINT \"PK_e94cda2f40a99b57e032a1a738b\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a012eaf5c87c65da1deb5fdbfa\" ON \"clip_note\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_ebe99317bbbe9968a0c6f579ad\" ON \"clip_note\" (\"clipId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_6fc0ec357d55a18646262fdfff\" ON \"clip_note\" (\"noteId\", \"clipId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"antenna_src_enum\" AS ENUM('home', 'all', 'list')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"antenna\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, \"src\" \"antenna_src_enum\" NOT NULL, \"userListId\" character varying(32), \"keywords\" jsonb NOT NULL DEFAULT '[]', \"withFile\" boolean NOT NULL, \"expression\" character varying(2048), \"notify\" boolean NOT NULL, \"hasNewNote\" boolean NOT NULL DEFAULT false, CONSTRAINT \"PK_c170b99775e1dccca947c9f2d5f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6446c571a0e8d0f05f01c78909\" ON \"antenna\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"antenna_note\" (\"id\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"antennaId\" character varying(32) NOT NULL, CONSTRAINT \"PK_fb28d94d0989a3872df19fd6ef8\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bd0397be22147e17210940e125\" ON \"antenna_note\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0d775946662d2575dfd2068a5f\" ON \"antenna_note\" (\"antennaId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_335a0bf3f904406f9ef3dd51c2\" ON \"antenna_note\" (\"noteId\", \"antennaId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP COLUMN \"geo\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"clip\" ADD CONSTRAINT \"FK_2b5ec6c574d6802c94c80313fb2\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"clip_note\" ADD CONSTRAINT \"FK_a012eaf5c87c65da1deb5fdbfa3\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"clip_note\" ADD CONSTRAINT \"FK_ebe99317bbbe9968a0c6f579adf\" FOREIGN KEY (\"clipId\") REFERENCES \"clip\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD CONSTRAINT \"FK_6446c571a0e8d0f05f01c789096\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD CONSTRAINT \"FK_709d7d32053d0dd7620f678eeb9\" FOREIGN KEY (\"userListId\") REFERENCES \"user_list\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna_note\" ADD CONSTRAINT \"FK_bd0397be22147e17210940e125b\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna_note\" ADD CONSTRAINT \"FK_0d775946662d2575dfd2068a5f5\" FOREIGN KEY (\"antennaId\") REFERENCES \"antenna\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" DROP COLUMN \"hasNewNote\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna_note\" ADD \"read\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_9937ea48d7ae97ffb4f3f063a4\" ON \"antenna_note\" (\"read\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"users\" character varying(1024) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"caseSensitive\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"antenna_src_enum\" RENAME TO \"antenna_src_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"antenna_src_enum\" AS ENUM('home', 'all', 'users', 'list')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ALTER COLUMN \"src\" TYPE \"antenna_src_enum\" USING \"src\"::\"text\"::\"antenna_src_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"antenna_src_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP CONSTRAINT \"FK_ec5c201576192ba8904c345c5cc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP COLUMN \"appId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"withReplies\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("TRUNCATE TABLE \"notification\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" DROP COLUMN \"type\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"notification_type_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'receiveFollowRequest', 'followRequestAccepted')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"type\" \"notification_type_enum\" NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" DROP COLUMN \"isMarkedAsClosed\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"isSuspended\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_34500da2e38ac393f7bb6b299c\" ON \"instance\" (\"isSuspended\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"twitter\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"twitterAccessToken\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"twitterAccessTokenSecret\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"twitterUserId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"twitterScreenName\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"github\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"githubAccessToken\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"githubId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"githubLogin\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discord\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordAccessToken\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordRefreshToken\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordExpiresDate\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordUsername\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"discordDiscriminator\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"integrations\" jsonb NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_NOTE_TAGS\" ON \"note\" USING gin (\"tags\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" RENAME COLUMN \"proxyAccount\" TO \"proxyAccountId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" DROP COLUMN \"proxyAccountId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"proxyAccountId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD CONSTRAINT \"FK_ab1bc0c1e209daa77b8e8d212ad\" FOREIGN KEY (\"proxyAccountId\") REFERENCES \"user\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_group_invitation\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"userGroupId\" character varying(32) NOT NULL, CONSTRAINT \"PK_160c63ec02bf23f6a5c5e8140d6\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bfbc6305547539369fe73eb144\" ON \"user_group_invitation\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5cc8c468090e129857e9fecce5\" ON \"user_group_invitation\" (\"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_e9793f65f504e5a31fbaedbf2f\" ON \"user_group_invitation\" (\"userId\", \"userGroupId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"userGroupInvitationId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"notification_type_enum\" RENAME TO \"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"notification_type_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'receiveFollowRequest', 'followRequestAccepted', 'groupInvited')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ALTER COLUMN \"type\" TYPE \"notification_type_enum\" USING \"type\"::\"text\"::\"notification_type_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"type\" IS 'The type of the Notification.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_invitation\" ADD CONSTRAINT \"FK_bfbc6305547539369fe73eb144a\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_group_invitation\" ADD CONSTRAINT \"FK_5cc8c468090e129857e9fecce5a\" FOREIGN KEY (\"userGroupId\") REFERENCES \"user_group\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_8fe87814e978053a53b1beb7e98\" FOREIGN KEY (\"userGroupInvitationId\") REFERENCES \"user_group_invitation\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"userGroupJoiningId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"antenna_src_enum\" RENAME TO \"antenna_src_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"antenna_src_enum\" AS ENUM('home', 'all', 'users', 'list', 'group')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ALTER COLUMN \"src\" TYPE \"antenna_src_enum\" USING \"src\"::\"text\"::\"antenna_src_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"antenna_src_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" DROP COLUMN \"users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"users\" character varying(1024) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD CONSTRAINT \"FK_ccbf5a8c0be4511133dcc50ddeb\" FOREIGN KEY (\"userGroupJoiningId\") REFERENCES \"user_group_joining\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_55720b33a61a7c806a8215b825\" ON \"drive_file\" (\"userId\", \"folderId\", \"id\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"promo_note\" (\"noteId\" character varying(32) NOT NULL, \"expiresAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, CONSTRAINT \"REL_e263909ca4fe5d57f8d4230dd5\" UNIQUE (\"noteId\"), CONSTRAINT \"PK_e263909ca4fe5d57f8d4230dd5c\" PRIMARY KEY (\"noteId\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_83f0862e9bae44af52ced7099e\" ON \"promo_note\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"promo_read\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, CONSTRAINT \"PK_61917c1541002422b703318b7c9\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_9657d55550c3d37bfafaf7d4b0\" ON \"promo_read\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_2882b8a1a07c7d281a98b6db16\" ON \"promo_read\" (\"userId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"promo_note\" ADD CONSTRAINT \"FK_e263909ca4fe5d57f8d4230dd5c\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"promo_read\" ADD CONSTRAINT \"FK_9657d55550c3d37bfafaf7d4b05\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"promo_read\" ADD CONSTRAINT \"FK_a46a1a603ecee695d7db26da5f4\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"injectFeaturedNote\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"antenna\" ADD \"excludeKeywords\" jsonb NOT NULL DEFAULT '[]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_reaction\" ALTER COLUMN \"reaction\" TYPE character varying(130)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"lastUsedAt\" TIMESTAMP WITH TIME ZONE DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"session\" character varying(128) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"name\" character varying(128) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"description\" character varying(512) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"iconUrl\" character varying(512) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"permission\" character varying(64) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD \"fetched\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" DROP CONSTRAINT \"FK_a3ff16c90cc87a82a0b5959e560\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"appId\" DROP NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"appId\" SET DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_bf3a053c07d9fb5d87317c56ee\" ON \"access_token\" (\"session\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ADD CONSTRAINT \"FK_a3ff16c90cc87a82a0b5959e560\" FOREIGN KEY (\"appId\") REFERENCES \"app\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"customBody\" character varying(2048)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"customHeader\" character varying(256)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"customIcon\" character varying(1024)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD \"appAccessTokenId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" DROP CONSTRAINT \"FK_3b4e96eec8d36a8bbb9d02aa710\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ALTER COLUMN \"notifierId\" DROP NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"notifierId\" IS 'The ID of sender user of the Notification.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"notification_type_enum\" RENAME TO \"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"notification_type_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'receiveFollowRequest', 'followRequestAccepted', 'groupInvited', 'app')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ALTER COLUMN \"type\" TYPE \"notification_type_enum\" USING \"type\"::\"text\"::\"notification_type_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"type\" IS 'The type of the Notification.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3b4e96eec8d36a8bbb9d02aa71\" ON \"notification\" (\"notifierId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_33f33cc8ef29d805a97ff4628b\" ON \"notification\" (\"type\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_080ab397c379af09b9d2169e5b\" ON \"notification\" (\"isRead\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_e22bf6bda77b6adc1fd9e75c8c\" ON \"notification\" (\"appAccessTokenId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_3b4e96eec8d36a8bbb9d02aa710\" FOREIGN KEY (\"notifierId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ADD CONSTRAINT \"FK_e22bf6bda77b6adc1fd9e75c8c9\" FOREIGN KEY (\"appAccessTokenId\") REFERENCES \"access_token\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD \"url\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageUseProxy\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_reaction\" ALTER COLUMN \"reaction\" TYPE character varying(260)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" ADD \"script\" character varying(16384) NOT NULL DEFAULT ''"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"enableHcaptcha\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"hcaptchaSiteKey\" character varying(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"hcaptchaSecretKey\" character varying(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"relay_status_enum\" AS ENUM('requesting', 'accepted', 'rejected')"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"relay\" (\"id\" character varying(32) NOT NULL, \"inbox\" character varying(512) NOT NULL, \"status\" \"relay_status_enum\" NOT NULL, CONSTRAINT \"PK_78ebc9cfddf4292633b7ba57aee\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0d9a1738f2cf7f3b1c3334dfab\" ON \"relay\" (\"inbox\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD \"blurhash\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"avatarColor\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"bannerColor\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"avatarBlurhash\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"bannerBlurhash\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"iconUrl\" character varying(256) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"muted_note\" (\"id\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, \"userId\" character varying(32) NOT NULL, CONSTRAINT \"PK_897e2eff1c0b9b64e55ca1418a4\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_70ab9786313d78e4201d81cdb8\" ON \"muted_note\" (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d8e07aa18c2d64e86201601aec\" ON \"muted_note\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a8c6bfd637d3f1d67a27c48e27\" ON \"muted_note\" (\"noteId\", \"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"enableWordMute\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"mutedWords\" jsonb NOT NULL DEFAULT '[]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3befe6f999c86aff06eb0257b4\" ON \"user_profile\" (\"enableWordMute\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muted_note\" ADD CONSTRAINT \"FK_70ab9786313d78e4201d81cdb89\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muted_note\" ADD CONSTRAINT \"FK_d8e07aa18c2d64e86201601aec1\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"muted_note_reason_enum\" AS ENUM('word', 'manual', 'spam', 'other')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muted_note\" ADD \"reason\" \"muted_note_reason_enum\" NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_636e977ff90b23676fb5624b25\" ON \"muted_note\" (\"reason\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"channel\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"lastNotedAt\" TIMESTAMP WITH TIME ZONE, \"userId\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, \"description\" character varying(2048), \"bannerId\" character varying(32), \"notesCount\" integer NOT NULL DEFAULT 0, \"usersCount\" integer NOT NULL DEFAULT 0, CONSTRAINT \"PK_590f33ee6ee7d76437acf362e39\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_71cb7b435b7c0d4843317e7e16\" ON \"channel\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_29ef80c6f13bcea998447fce43\" ON \"channel\" (\"lastNotedAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_823bae55bd81b3be6e05cff438\" ON \"channel\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0f58c11241e649d2a638a8de94\" ON \"channel\" (\"notesCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_094b86cd36bb805d1aa1e8cc9a\" ON \"channel\" (\"usersCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"channel_following\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"followeeId\" character varying(32) NOT NULL, \"followerId\" character varying(32) NOT NULL, CONSTRAINT \"PK_8b104be7f7415113f2a02cd5bdd\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_11e71f2511589dcc8a4d3214f9\" ON \"channel_following\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0e43068c3f92cab197c3d3cd86\" ON \"channel_following\" (\"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6d8084ec9496e7334a4602707e\" ON \"channel_following\" (\"followerId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_2e230dd45a10e671d781d99f3e\" ON \"channel_following\" (\"followerId\", \"followeeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"channel_note_pining\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"channelId\" character varying(32) NOT NULL, \"noteId\" character varying(32) NOT NULL, CONSTRAINT \"PK_44f7474496bcf2e4b741681146d\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8125f950afd3093acb10d2db8a\" ON \"channel_note_pining\" (\"channelId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_f36fed37d6d4cdcc68c803cd9c\" ON \"channel_note_pining\" (\"channelId\", \"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD \"channelId\" character varying(32) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f22169eb10657bded6d875ac8f\" ON \"note\" (\"channelId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel\" ADD CONSTRAINT \"FK_823bae55bd81b3be6e05cff4383\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel\" ADD CONSTRAINT \"FK_999da2bcc7efadbfe0e92d3bc19\" FOREIGN KEY (\"bannerId\") REFERENCES \"drive_file\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD CONSTRAINT \"FK_f22169eb10657bded6d875ac8f9\" FOREIGN KEY (\"channelId\") REFERENCES \"channel\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" ADD CONSTRAINT \"FK_0e43068c3f92cab197c3d3cd86e\" FOREIGN KEY (\"followeeId\") REFERENCES \"channel\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" ADD CONSTRAINT \"FK_6d8084ec9496e7334a4602707e1\" FOREIGN KEY (\"followerId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_note_pining\" ADD CONSTRAINT \"FK_8125f950afd3093acb10d2db8a8\" FOREIGN KEY (\"channelId\") REFERENCES \"channel\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_note_pining\" ADD CONSTRAINT \"FK_10b19ef67d297ea9de325cd4502\" FOREIGN KEY (\"noteId\") REFERENCES \"note\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" ADD \"readCursor\" TIMESTAMP WITH TIME ZONE NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageSetPublicRead\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"user_profile_includingnotificationtypes_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'receiveFollowRequest', 'followRequestAccepted', 'groupInvited', 'app')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"includingNotificationTypes\" \"user_profile_includingnotificationtypes_enum\" array"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a7eba67f8b3fa27271e85d2e26\" ON \"drive_file\" (\"isSensitive\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("TRUNCATE TABLE \"note_unread\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel_following\" DROP COLUMN \"readCursor\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" ADD \"isMentioned\" boolean NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_unread\" ADD \"noteChannelId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_25b1dd384bec391b07b74b861c\" ON \"note_unread\" (\"isMentioned\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_89a29c9237b8c3b6b3cbb4cb30\" ON \"note_unread\" (\"isSpecified\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_29e8c1d579af54d4232939f994\" ON \"note_unread\" (\"noteUserId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_6a57f051d82c6d4036c141e107\" ON \"note_unread\" (\"noteChannelId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f22169eb10657bded6d875ac8f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_note_on_channelId_and_id_desc\" ON \"note\" (\"channelId\", \"id\" desc)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"includingNotificationTypes\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"user_profile_includingnotificationtypes_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"user_profile_mutingnotificationtypes_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'receiveFollowRequest', 'followRequestAccepted', 'groupInvited', 'app')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"mutingNotificationTypes\" \"user_profile_mutingnotificationtypes_enum\" array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" DROP CONSTRAINT \"FK_d049123c413e68ca52abe734203\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_d049123c413e68ca52abe73420\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_5cd442c3b2e74fdd99dae20243\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" RENAME COLUMN \"userId\" TO \"targetUserId\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"assigneeId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"resolved\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" DROP COLUMN \"comment\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"comment\" character varying(2048) NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2b15aaf4a0dc5be3499af7ab6a\" ON \"abuse_user_report\" (\"resolved\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD CONSTRAINT \"FK_08b883dd5fdd6f9c4c1572b36de\" FOREIGN KEY (\"assigneeId\") REFERENCES \"user\"(\"id\") ON DELETE SET NULL ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"targetUserHost\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"reporterHost\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4ebbf7f93cdc10e8d1ef2fc6cd\" ON \"abuse_user_report\" (\"targetUserHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f8d8b93740ad12c4ce8213a199\" ON \"abuse_user_report\" (\"reporterHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"themeColor\" character varying(64) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ADD \"faviconUrl\" character varying(256) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" DROP COLUMN \"autoWatch\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"clip\" ADD \"description\" character varying(2048) DEFAULT null"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"log\".\"createdAt\" IS 'The created date of the Log.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_folder\".\"createdAt\" IS 'The created date of the DriveFolder.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_folder\".\"name\" IS 'The name of the DriveFolder.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_folder\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_folder\".\"parentId\" IS 'The parent folder ID. If null, it means the DriveFolder is located in root.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"createdAt\" IS 'The created date of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"userHost\" IS 'The host of owner. It will be null if the user in local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"md5\" IS 'The MD5 hash of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"name\" IS 'The file name of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"type\" IS 'The content type (MIME) of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"size\" IS 'The file size (bytes) of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"comment\" IS 'The comment of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"blurhash\" IS 'The BlurHash string.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"properties\" IS 'The any properties of the DriveFile. For example, it includes image width/height.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"url\" IS 'The URL of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"thumbnailUrl\" IS 'The URL of the thumbnail of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"webpublicUrl\" IS 'The URL of the webpublic of the DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"uri\" IS 'The URI of the DriveFile. it will be null when the DriveFile is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"folderId\" IS 'The parent folder ID. If null, it means the DriveFile is located in root.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"isSensitive\" IS 'Whether the DriveFile is NSFW.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"drive_file\".\"isLink\" IS 'Whether the DriveFile is direct link to remote server.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"createdAt\" IS 'The created date of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"updatedAt\" IS 'The updated date of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"username\" IS 'The username of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"usernameLower\" IS 'The username (lowercased) of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"name\" IS 'The name of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"followersCount\" IS 'The count of followers.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"followingCount\" IS 'The count of following.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"notesCount\" IS 'The count of notes.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"avatarId\" IS 'The ID of avatar DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"bannerId\" IS 'The ID of banner DriveFile.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isSuspended\" IS 'Whether the User is suspended.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isSilenced\" IS 'Whether the User is silenced.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isLocked\" IS 'Whether the User is locked.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isBot\" IS 'Whether the User is a bot.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isCat\" IS 'Whether the User is a cat.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isAdmin\" IS 'Whether the User is the admin.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isModerator\" IS 'Whether the User is a moderator.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"host\" IS 'The host of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"inbox\" IS 'The inbox URL of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"sharedInbox\" IS 'The sharedInbox URL of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"featured\" IS 'The featured URL of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"uri\" IS 'The URI of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"token\" IS 'The native access token of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"createdAt\" IS 'The created date of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"secret\" IS 'The secret key of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"name\" IS 'The name of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"description\" IS 'The description of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"permission\" IS 'The permission of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"app\".\"callbackUrl\" IS 'The callbackUrl of the App.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"createdAt\" IS 'The created date of the AccessToken.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"lastUsedAt\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"session\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"appId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"name\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"description\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"access_token\".\"iconUrl\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"createdAt\" IS 'The created date of the Channel.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"name\" IS 'The name of the Channel.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"description\" IS 'The description of the Channel.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"bannerId\" IS 'The ID of banner Channel.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"notesCount\" IS 'The count of notes.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel\".\"usersCount\" IS 'The count of users.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"createdAt\" IS 'The created date of the Note.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"replyId\" IS 'The ID of reply target.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"renoteId\" IS 'The ID of renote target.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"userId\" IS 'The ID of author.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"uri\" IS 'The URI of a note. it will be null when the note is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"url\" IS 'The human readable url of a note. it will be null when the note is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"channelId\" IS 'The ID of source channel.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"userHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"replyUserId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"replyUserHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"renoteUserId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note\".\"renoteUserHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"poll_vote\".\"createdAt\" IS 'The created date of the PollVote.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_reaction\".\"createdAt\" IS 'The created date of the NoteReaction.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_watching\".\"createdAt\" IS 'The created date of the NoteWatching.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_watching\".\"userId\" IS 'The watcher ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_watching\".\"noteId\" IS 'The target Note ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_watching\".\"noteUserId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_unread\".\"noteUserId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_unread\".\"noteChannelId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"createdAt\" IS 'The created date of the FollowRequest.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followeeId\" IS 'The followee user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followerId\" IS 'The follower user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"requestId\" IS 'id of Follow Activity.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followerHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followerInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followerSharedInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followeeHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followeeInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"follow_request\".\"followeeSharedInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group\".\"createdAt\" IS 'The created date of the UserGroup.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group\".\"userId\" IS 'The ID of owner.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_invitation\".\"createdAt\" IS 'The created date of the UserGroupInvitation.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_invitation\".\"userId\" IS 'The user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_invitation\".\"userGroupId\" IS 'The group ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"createdAt\" IS 'The created date of the Notification.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"notifieeId\" IS 'The ID of recipient user of the Notification.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"notification\".\"isRead\" IS 'Whether the Notification is read.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"meta\".\"localDriveCapacityMb\" IS 'Drive capacity of a local user (MB)'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"meta\".\"remoteDriveCapacityMb\" IS 'Drive capacity of a remote user (MB)'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"meta\".\"maxNoteTextLength\" IS 'Max allowed note text length in characters'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"createdAt\" IS 'The created date of the Following.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followeeId\" IS 'The followee user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followerId\" IS 'The follower user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followerHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followerInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followerSharedInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followeeHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followeeInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"following\".\"followeeSharedInbox\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"caughtAt\" IS 'The caught date of the Instance.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"host\" IS 'The host of the Instance.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"usersCount\" IS 'The count of the users of the Instance.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"notesCount\" IS 'The count of the notes of the Instance.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"softwareName\" IS 'The software of the Instance.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"softwareVersion\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"openRegistrations\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"name\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"description\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"maintainerName\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"maintainerEmail\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"iconUrl\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"faviconUrl\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"instance\".\"themeColor\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muting\".\"createdAt\" IS 'The created date of the Muting.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muting\".\"muteeId\" IS 'The mutee user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muting\".\"muterId\" IS 'The muter user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"blocking\".\"createdAt\" IS 'The created date of the Blocking.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"blocking\".\"blockeeId\" IS 'The blockee user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"blocking\".\"blockerId\" IS 'The blocker user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list\".\"createdAt\" IS 'The created date of the UserList.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list\".\"name\" IS 'The name of the UserList.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list_joining\".\"createdAt\" IS 'The created date of the UserListJoining.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list_joining\".\"userId\" IS 'The user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_list_joining\".\"userListId\" IS 'The list ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_joining\".\"createdAt\" IS 'The created date of the UserGroupJoining.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_joining\".\"userId\" IS 'The user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_group_joining\".\"userGroupId\" IS 'The group ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"note_favorite\".\"createdAt\" IS 'The created date of the NoteFavorite.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"abuse_user_report\".\"createdAt\" IS 'The created date of the AbuseUserReport.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"abuse_user_report\".\"targetUserHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"abuse_user_report\".\"reporterHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"messaging_message\".\"createdAt\" IS 'The created date of the MessagingMessage.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"messaging_message\".\"userId\" IS 'The sender user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"messaging_message\".\"groupId\" IS 'The recipient group ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"signin\".\"createdAt\" IS 'The created date of the Signin.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"auth_session\".\"createdAt\" IS 'The created date of the AuthSession.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"reversi_game\".\"createdAt\" IS 'The created date of the ReversiGame.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"reversi_game\".\"startedAt\" IS 'The started date of the ReversiGame.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"reversi_game\".\"form1\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"reversi_game\".\"form2\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"reversi_matching\".\"createdAt\" IS 'The created date of the ReversiMatching.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_note_pining\".\"createdAt\" IS 'The created date of the UserNotePinings.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"poll\".\"noteId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"poll\".\"noteVisibility\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"poll\".\"userId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"poll\".\"userHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_keypair\".\"userId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_publickey\".\"userId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"page\".\"createdAt\" IS 'The created date of the Page.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"page\".\"updatedAt\" IS 'The updated date of the Page.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"page\".\"userId\" IS 'The ID of author.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"userId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"location\" IS 'The location of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"birthday\" IS 'The birthday (YYYY-MM-DD) of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"description\" IS 'The description (bio) of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"url\" IS 'Remote URL of the user.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"email\" IS 'The email address of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"password\" IS 'The password hash of the User. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"clientData\" IS 'The client-specific data of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"room\" IS 'The room data of the User.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"userHost\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_security_key\".\"id\" IS 'Variable-length id given to navigator.credentials.get()'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_security_key\".\"publicKey\" IS 'Variable-length public key used to verify attestations (hex-encoded).'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_security_key\".\"lastUsed\" IS 'The date of the last time the UserSecurityKey was successfully validated.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_security_key\".\"name\" IS 'User-defined name for this key'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"attestation_challenge\".\"challenge\" IS 'Hex-encoded sha256 hash of the challenge.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"attestation_challenge\".\"createdAt\" IS 'The date challenge was created for expiry purposes.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"attestation_challenge\".\"registrationChallenge\" IS 'Indicates that the challenge is only for registration purposes if true to prevent the challenge for being used as authentication.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"moderation_log\".\"createdAt\" IS 'The created date of the ModerationLog.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"announcement\".\"createdAt\" IS 'The created date of the Announcement.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"announcement\".\"updatedAt\" IS 'The updated date of the Announcement.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"announcement_read\".\"createdAt\" IS 'The created date of the AnnouncementRead.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip\".\"createdAt\" IS 'The created date of the Clip.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip\".\"name\" IS 'The name of the Clip.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip\".\"description\" IS 'The description of the Clip.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip_note\".\"noteId\" IS 'The note ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"clip_note\".\"clipId\" IS 'The clip ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"antenna\".\"createdAt\" IS 'The created date of the Antenna.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"antenna\".\"userId\" IS 'The owner ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"antenna\".\"name\" IS 'The name of the Antenna.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"antenna_note\".\"noteId\" IS 'The note ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"antenna_note\".\"antennaId\" IS 'The antenna ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"promo_note\".\"noteId\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"promo_note\".\"userId\" IS '[Denormalized]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"promo_read\".\"createdAt\" IS 'The created date of the PromoRead.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muted_note\".\"noteId\" IS 'The note ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muted_note\".\"userId\" IS 'The user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"muted_note\".\"reason\" IS 'The reason of the MutedNote.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel_following\".\"createdAt\" IS 'The created date of the ChannelFollowing.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel_following\".\"followeeId\" IS 'The followee channel ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel_following\".\"followerId\" IS 'The follower user ID.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"channel_note_pining\".\"createdAt\" IS 'The created date of the ChannelNotePining.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"pinnedPages\" character varying(512) array NOT NULL DEFAULT '{\"/featured\", \"/channels\", \"/explore\", \"/pages\", \"/about-misskey\"}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"backgroundImageUrl\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"logoImageUrl\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"noCrawle\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"noCrawle\" IS 'Whether reject index by crawler.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"pinnedClipId\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"isExplorable\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isExplorable\" IS 'Whether the User is explorable.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d5a1b83c7cab66f167e6888188\" ON \"user\" (\"isExplorable\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"registry_item\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"updatedAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"key\" character varying(1024) NOT NULL, \"scope\" character varying(1024) array NOT NULL DEFAULT '{}'::varchar[], \"domain\" character varying(512), CONSTRAINT \"PK_64b3f7e6008b4d89b826cd3af95\" PRIMARY KEY (\"id\")); COMMENT ON COLUMN \"registry_item\".\"createdAt\" IS 'The created date of the RegistryItem.'; COMMENT ON COLUMN \"registry_item\".\"updatedAt\" IS 'The updated date of the RegistryItem.'; COMMENT ON COLUMN \"registry_item\".\"userId\" IS 'The owner ID.'; COMMENT ON COLUMN \"registry_item\".\"key\" IS 'The key of the RegistryItem.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_fb9d21ba0abb83223263df6bcb\" ON \"registry_item\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_22baca135bb8a3ea1a83d13df3\" ON \"registry_item\" (\"scope\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_0a72bdfcdb97c0eca11fe7ecad\" ON \"registry_item\" (\"domain\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"registry_item\" ADD CONSTRAINT \"FK_fb9d21ba0abb83223263df6bcb3\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"registry_item\" ADD \"value\" jsonb NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"registry_item\".\"value\" IS 'The value of the RegistryItem.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"registry_item\" ALTER COLUMN \"value\" DROP NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"followersUri\" varchar(512) DEFAULT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"followersUri\" IS 'The URI of the user Follower Collection. It will be null if the origin of the user is local.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"gallery_post\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"updatedAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"title\" character varying(256) NOT NULL, \"description\" character varying(2048), \"userId\" character varying(32) NOT NULL, \"fileIds\" character varying(32) array NOT NULL DEFAULT '{}'::varchar[], \"isSensitive\" boolean NOT NULL DEFAULT false, \"likedCount\" integer NOT NULL DEFAULT '0', \"tags\" character varying(128) array NOT NULL DEFAULT '{}'::varchar[], CONSTRAINT \"PK_8e90d7b6015f2c4518881b14753\" PRIMARY KEY (\"id\")); COMMENT ON COLUMN \"gallery_post\".\"createdAt\" IS 'The created date of the GalleryPost.'; COMMENT ON COLUMN \"gallery_post\".\"updatedAt\" IS 'The updated date of the GalleryPost.'; COMMENT ON COLUMN \"gallery_post\".\"userId\" IS 'The ID of author.'; COMMENT ON COLUMN \"gallery_post\".\"isSensitive\" IS 'Whether the post is sensitive.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8f1a239bd077c8864a20c62c2c\" ON \"gallery_post\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f631d37835adb04792e361807c\" ON \"gallery_post\" (\"updatedAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_985b836dddd8615e432d7043dd\" ON \"gallery_post\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_3ca50563facd913c425e7a89ee\" ON \"gallery_post\" (\"fileIds\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f2d744d9a14d0dfb8b96cb7fc5\" ON \"gallery_post\" (\"isSensitive\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_1a165c68a49d08f11caffbd206\" ON \"gallery_post\" (\"likedCount\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_05cca34b985d1b8edc1d1e28df\" ON \"gallery_post\" (\"tags\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"gallery_like\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"postId\" character varying(32) NOT NULL, CONSTRAINT \"PK_853ab02be39b8de45cd720cc15f\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8fd5215095473061855ceb948c\" ON \"gallery_like\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_df1b5f4099e99fb0bc5eae53b6\" ON \"gallery_like\" (\"userId\", \"postId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"gallery_post\" ADD CONSTRAINT \"FK_985b836dddd8615e432d7043ddb\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"gallery_like\" ADD CONSTRAINT \"FK_8fd5215095473061855ceb948cf\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"gallery_like\" ADD CONSTRAINT \"FK_b1cb568bfe569e47b7051699fc8\" FOREIGN KEY (\"postId\") REFERENCES \"gallery_post\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"objectStorageS3ForcePathStyle\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"receiveAnnouncementEmail\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"emailNotificationTypes\" jsonb NOT NULL DEFAULT '[\"follow\",\"receiveFollowRequest\",\"groupInvited\"]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"lang\" character varying(32)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"driveUsage\" TYPE bigint"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__active_users\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__drive\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__federation\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__hashtag\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__instance\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__network\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__notes\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_drive\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_following\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_notes\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_reaction\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__test\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__test_grouped\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__test_unique\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__users\" WHERE \"span\" = 'day'"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_15e91a03aeeac9dbccdf43fc06\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_20f57cc8f142c131340ee16742\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_c26e2c1cbb6e911e0554b27416\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_3fa0d0f17ca72e3dc80999a032\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_6e1df243476e20cbf86572ecc0\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_06690fc959f1c9fdaf21928222\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_e447064455928cf627590ef527\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_2d416e6af791a82e338c79d480\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_e9cd07672b37d8966cf3709283\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_fcc181fb8283009c61cc4083ef\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_49975586f50ed7b800fdd88fbd\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_6d6f156ceefc6bc5f273a0e370\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_c12f0af4a66cdd30c2287ce8aa\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_d0a4f79af5a97b08f37b547197\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f5448d9633cff74208d850aabe\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f8dd01baeded2ffa833e0a610a\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_08fac0eb3b11f04c200c0b40dd\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_9ff6944f01acb756fdc92d7563\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_e69096589f11e3baa98ddd64d0\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0c9a159c5082cbeef3ca6706b5\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_924fc196c80ca24bae01dd37e4\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_328f259961e60c4fa0bfcf55ca\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_42ea9381f0fda8dfe0fa1c8b53\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f2aeafde2ae6fbad38e857631b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f92dd6d03f8d994f29987f6214\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_57b5458d0d3d6d1e7f13d4e57f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_4db3b84c7be0d3464714f3e0b1\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_8d2cbbc8114d90d19b44d626b6\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_046feeb12e9ef5f783f409866a\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f68a5ab958f9f5fa17a32ac23b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_65633a106bce43fc7c5c30a5c7\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_edeb73c09c3143a81bcb34d569\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_e316f01a6d24eb31db27f88262\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_2be7ec6cebddc14dc11e206686\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_a5133470f4825902e170328ca5\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_84e661abb7bd1e51b690d4b017\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_5c73bf61da4f6e6f15bae88ed1\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_d70c86baedc68326be11f9c0ce\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_66e1e1ecd2f29e57778af35b59\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_92255988735563f0fe4aba1f05\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_c5870993e25c3d5771f91f5003\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f170de677ea75ad4533de2723e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_7c184198ecf66a8d3ecb253ab3\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f091abb24193d50c653c6b77fc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_a770a57c70e668cc61590c9161\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__active_users_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___local_count\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___remote_count\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__drive_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__federation_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__hashtag_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"___local_count\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"___remote_count\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__instance_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__network_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__notes_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__per_user_drive_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__per_user_following_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__per_user_notes_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__per_user_reaction_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_grouped\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__test_grouped_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_grouped\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_unique\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__test_unique_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_unique\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_unique\" DROP COLUMN \"___foo\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__test_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" DROP COLUMN \"span\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"__chart__users_span_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" DROP COLUMN \"unique\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___local_users\" character varying array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___remote_users\" character varying array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"___local_users\" character varying array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"___remote_users\" character varying array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__test_unique\" ADD \"___foo\" character varying array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"lastActiveDate\" TIMESTAMP WITH TIME ZONE DEFAULT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_seoignmeoprigmkpodgrjmkpormg\" ON \"user\" (\"lastActiveDate\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"hideOnlineStatus\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"password_reset_request\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"token\" character varying(256) NOT NULL, \"userId\" character varying(32) NOT NULL, CONSTRAINT \"PK_fcf4b02eae1403a2edaf87fd074\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0b575fa9a4cfe638a925949285\" ON \"password_reset_request\" (\"token\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4bb7fd4a34492ae0e6cc8d30ac\" ON \"password_reset_request\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"password_reset_request\" ADD CONSTRAINT \"FK_4bb7fd4a34492ae0e6cc8d30ac8\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"ad\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"expiresAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"place\" character varying(32) NOT NULL, \"priority\" character varying(32) NOT NULL, \"url\" character varying(1024) NOT NULL, \"imageUrl\" character varying(1024) NOT NULL, \"memo\" character varying(8192) NOT NULL, CONSTRAINT \"PK_0193d5ef09746e88e9ea92c634d\" PRIMARY KEY (\"id\")); COMMENT ON COLUMN \"ad\".\"createdAt\" IS 'The created date of the Ad.'; COMMENT ON COLUMN \"ad\".\"expiresAt\" IS 'The expired date of the Ad.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_1129c2ef687fc272df040bafaa\" ON \"ad\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_2da24ce20ad209f1d9dc032457\" ON \"ad\" (\"expiresAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"ad\" ADD \"ratio\" integer NOT NULL DEFAULT '1'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_NOTE_MENTIONS\" ON \"note\" USING gin (\"mentions\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_NOTE_VISIBLE_USER_IDS\" ON \"note\" USING gin (\"visibleUserIds\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ALTER COLUMN \"description\" TYPE character varying(2048)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"log\" ALTER COLUMN \"message\" TYPE character varying(2048)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"avatarUrl\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"bannerUrl\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"avatarBlurhash\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" DROP COLUMN \"bannerBlurhash\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" DROP COLUMN \"proxyRemoteFiles\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__active_users\" a USING \"__chart__active_users\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__drive\" a USING \"__chart__drive\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__federation\" a USING \"__chart__federation\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__hashtag\" a USING \"__chart__hashtag\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__instance\" a USING \"__chart__instance\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__network\" a USING \"__chart__network\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__notes\" a USING \"__chart__notes\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_drive\" a USING \"__chart__per_user_drive\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_following\" a USING \"__chart__per_user_following\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_notes\" a USING \"__chart__per_user_notes\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_reaction\" a USING \"__chart__per_user_reaction\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__test_grouped\" a USING \"__chart__test_grouped\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__test_unique\" a USING \"__chart__test_unique\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__users\" a USING \"__chart__users\" b WHERE a.id < b.id AND ((a.group IS NULL AND b.group IS NULL) OR a.group = b.group) AND a.date = b.date;"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0ad37b7ef50f4ddc84363d7ccc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_00ed5f86db1f7efafb1978bf21\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_9a3ed15a30ab7e3a37702e6e08\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_13565815f618a1ff53886c5b28\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_7a170f67425e62a8fabb76c872\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_3313d7288855ec105b5bbf6c21\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_36cb699c49580d4e6c2e6159f9\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_76e87c7bfc5d925fcbba405d84\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_dd907becf76104e4b656659e6b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_07747a1038c05f532a718fe1de\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_99a7d2faaef84a6f728d714ad6\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_25a97c02003338124b2b75fdbc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_6b8f34a1a64b06014b6fb66824\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_da8a46ba84ca1d8bb5a29bfb63\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_39ee857ab2f23493037c6b6631\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_a1efd3e0048a5f2793a47360dc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_7b5da130992ec9df96712d4290\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0a905b992fecd2b5c3fb98759e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_42eb716a37d381cdf566192b2b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_7036f2957151588b813185c794\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f09d543e3acb16c5976bdb31fa\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_5f86db6492274e07c1a3cdf286\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_e496ca8096d28f6b9b509264dc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_30bf67687f483ace115c5ca642\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_7af07790712aa3438ff6773f3b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_4b3593098b6edc9c5afe36b18b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_b77d4dd9562c3a899d9a286fcd\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_84234bd1abb873f07329681c83\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_55bf20f366979f2436de99206b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_5048e9daccbbbc6d567bb142d3\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_f7bf4c62059764c2c2bb40fdab\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_8cf3156fd7a6b15c43459c6e3b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_229a41ad465f9205f1f5703291\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_0c641990ecf47d2545df4edb75\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_234dff3c0b56a6150b95431ab9\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_b14489029e4b3aaf4bba5fb524\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_437bab3c6061d90f6bb65fd2cc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_bbfa573a8181018851ed0b6357\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_a0cd75442dd10d0643a17c4a49\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_b070a906db04b44c67c6c2144d\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_d41cce6aee1a50bfc062038f9b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_a319e5dbf47e8a17497623beae\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_845254b3eaf708ae8a6cac3026\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_ed9b95919c672a13008e9487ee\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"IDX_337e9599f278bd7537fe30876f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_9a3ed15a30ab7e3a37702e6e08\" ON \"__chart__active_users\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_60c5c6e7e538c09aa274ecd1cf\" ON \"__chart__active_users\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_3313d7288855ec105b5bbf6c21\" ON \"__chart__drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_ceab80a6729f8e2e6f5b8a1a3d\" ON \"__chart__drive\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_dd907becf76104e4b656659e6b\" ON \"__chart__federation\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_eddfed8fb40305a04c6f941050\" ON \"__chart__federation\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_25a97c02003338124b2b75fdbc\" ON \"__chart__hashtag\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_53a3604b939e2b479eb2cfaac8\" ON \"__chart__hashtag\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_39ee857ab2f23493037c6b6631\" ON \"__chart__instance\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_8111b817b9818c04d7eb8475b1\" ON \"__chart__instance\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0a905b992fecd2b5c3fb98759e\" ON \"__chart__network\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_2082327b2699ce924fa654afc5\" ON \"__chart__network\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_f09d543e3acb16c5976bdb31fa\" ON \"__chart__notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_e60c358aaced5aab8900a4af31\" ON \"__chart__notes\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_30bf67687f483ace115c5ca642\" ON \"__chart__per_user_drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a9a806d466b314f253a1a611c4\" ON \"__chart__per_user_drive\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_b77d4dd9562c3a899d9a286fcd\" ON \"__chart__per_user_following\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_dabbb38a51ab86ee3cab291326\" ON \"__chart__per_user_following\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_5048e9daccbbbc6d567bb142d3\" ON \"__chart__per_user_notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_583a157ed0cf0ed1b5ec2a833f\" ON \"__chart__per_user_notes\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_229a41ad465f9205f1f5703291\" ON \"__chart__per_user_reaction\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_3b7697a96f522d0478972e6d6f\" ON \"__chart__per_user_reaction\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_b14489029e4b3aaf4bba5fb524\" ON \"__chart__test_grouped\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_da522b4008a9f5d7743b87ad55\" ON \"__chart__test_grouped\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a0cd75442dd10d0643a17c4a49\" ON \"__chart__test_unique\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_16effb2e888f6763673b579f80\" ON \"__chart__test_unique\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a319e5dbf47e8a17497623beae\" ON \"__chart__test\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_dab383a36f3c9db4a0c9b02cf3\" ON \"__chart__test\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_337e9599f278bd7537fe30876f\" ON \"__chart__users\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_66feba81e1795d176d06c0b1e6\" ON \"__chart__users\" (\"date\") WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"deeplAuthKey\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"channel\" ALTER COLUMN \"userId\" DROP NOT NULL;"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"isDeleted\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"isDeleted\" IS 'Whether the User is deleted.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"deeplIsPro\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"showTimelineReplies\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"showTimelineReplies\" IS 'Whether to show users replying to other users in the timeline.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"mutedInstances\" jsonb NOT NULL DEFAULT '[]'"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user_profile\".\"mutedInstances\" IS 'List of instances muted by the user.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"emailRequiredForSignup\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_pending\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"code\" character varying(128) NOT NULL, \"username\" character varying(128) NOT NULL, \"email\" character varying(128) NOT NULL, \"password\" character varying(128) NOT NULL, CONSTRAINT \"PK_d4c84e013c98ec02d19b8fbbafa\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_4e5c4c99175638ec0761714ab0\" ON \"user_pending\" (\"code\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"publicReactions\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TABLE \"log\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"note_thread_muting\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"threadId\" character varying(256) NOT NULL, CONSTRAINT \"PK_ec5936d94d1a0369646d12a3a47\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_29c11c7deb06615076f8c95b80\" ON \"note_thread_muting\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c426394644267453e76f036926\" ON \"note_thread_muting\" (\"threadId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_ae7aab18a2641d3e5f25e0c4ea\" ON \"note_thread_muting\" (\"userId\", \"threadId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD \"threadId\" character varying(256)"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_d4ebdef929896d6dc4a3c5bb48\" ON \"note\" (\"threadId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note_thread_muting\" ADD CONSTRAINT \"FK_29c11c7deb06615076f8c95b80a\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"public\".\"user_profile_ffvisibility_enum\" AS ENUM('public', 'followers', 'private')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"ffVisibility\" \"public\".\"user_profile_ffvisibility_enum\" NOT NULL DEFAULT 'public'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" DROP COLUMN \"viaMobile\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD \"forwarded\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"__chart__per_user_drive\" WHERE \"group\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_dd907becf76104e4b656659e6b\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_eddfed8fb40305a04c6f941050\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_f09d543e3acb16c5976bdb31fa\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_e60c358aaced5aab8900a4af31\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_337e9599f278bd7537fe30876f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_66feba81e1795d176d06c0b1e6\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_0a905b992fecd2b5c3fb98759e\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_2082327b2699ce924fa654afc5\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_9a3ed15a30ab7e3a37702e6e08\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_60c5c6e7e538c09aa274ecd1cf\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_8111b817b9818c04d7eb8475b1\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_583a157ed0cf0ed1b5ec2a833f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_3313d7288855ec105b5bbf6c21\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_ceab80a6729f8e2e6f5b8a1a3d\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_3b7697a96f522d0478972e6d6f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_53a3604b939e2b479eb2cfaac8\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_dabbb38a51ab86ee3cab291326\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_a9a806d466b314f253a1a611c4\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__federation\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___instance_total\" bigint NOT NULL, \"___instance_inc\" bigint NOT NULL, \"___instance_dec\" bigint NOT NULL, CONSTRAINT \"UQ_617a8fe225a6e701d89e02d2c74\" UNIQUE (\"date\"), CONSTRAINT \"PK_7ca721c769f31698e0e1331e8e6\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_617a8fe225a6e701d89e02d2c7\" ON \"__chart_day__federation\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__notes\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___local_total\" bigint NOT NULL, \"___local_inc\" bigint NOT NULL, \"___local_dec\" bigint NOT NULL, \"___local_diffs_normal\" bigint NOT NULL, \"___local_diffs_reply\" bigint NOT NULL, \"___local_diffs_renote\" bigint NOT NULL, \"___remote_total\" bigint NOT NULL, \"___remote_inc\" bigint NOT NULL, \"___remote_dec\" bigint NOT NULL, \"___remote_diffs_normal\" bigint NOT NULL, \"___remote_diffs_reply\" bigint NOT NULL, \"___remote_diffs_renote\" bigint NOT NULL, CONSTRAINT \"UQ_1a527b423ad0858a1af5a056d43\" UNIQUE (\"date\"), CONSTRAINT \"PK_1fa4139e1f338272b758d05e090\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_1a527b423ad0858a1af5a056d4\" ON \"__chart_day__notes\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__users\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___local_total\" bigint NOT NULL, \"___local_inc\" bigint NOT NULL, \"___local_dec\" bigint NOT NULL, \"___remote_total\" bigint NOT NULL, \"___remote_inc\" bigint NOT NULL, \"___remote_dec\" bigint NOT NULL, CONSTRAINT \"UQ_cad6e07c20037f31cdba8a350c3\" UNIQUE (\"date\"), CONSTRAINT \"PK_d7f7185abb9851f70c4726c54bd\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_cad6e07c20037f31cdba8a350c\" ON \"__chart_day__users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__network\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___incomingRequests\" bigint NOT NULL, \"___outgoingRequests\" bigint NOT NULL, \"___totalTime\" bigint NOT NULL, \"___incomingBytes\" bigint NOT NULL, \"___outgoingBytes\" bigint NOT NULL, CONSTRAINT \"UQ_8bfa548c2b31f9e07db113773ee\" UNIQUE (\"date\"), CONSTRAINT \"PK_cac499d6f471042dfed1e7e0132\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_8bfa548c2b31f9e07db113773e\" ON \"__chart_day__network\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__active_users\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___local_users\" character varying array NOT NULL, \"___remote_users\" character varying array NOT NULL, CONSTRAINT \"UQ_d5954f3df5e5e3bdfc3c03f3906\" UNIQUE (\"date\"), CONSTRAINT \"PK_b1790489b14f005ae8f404f5795\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d5954f3df5e5e3bdfc3c03f390\" ON \"__chart_day__active_users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__instance\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___requests_failed\" bigint NOT NULL, \"___requests_succeeded\" bigint NOT NULL, \"___requests_received\" bigint NOT NULL, \"___notes_total\" bigint NOT NULL, \"___notes_inc\" bigint NOT NULL, \"___notes_dec\" bigint NOT NULL, \"___notes_diffs_normal\" bigint NOT NULL, \"___notes_diffs_reply\" bigint NOT NULL, \"___notes_diffs_renote\" bigint NOT NULL, \"___users_total\" bigint NOT NULL, \"___users_inc\" bigint NOT NULL, \"___users_dec\" bigint NOT NULL, \"___following_total\" bigint NOT NULL, \"___following_inc\" bigint NOT NULL, \"___following_dec\" bigint NOT NULL, \"___followers_total\" bigint NOT NULL, \"___followers_inc\" bigint NOT NULL, \"___followers_dec\" bigint NOT NULL, \"___drive_totalFiles\" bigint NOT NULL, \"___drive_totalUsage\" bigint NOT NULL, \"___drive_incFiles\" bigint NOT NULL, \"___drive_incUsage\" bigint NOT NULL, \"___drive_decFiles\" bigint NOT NULL, \"___drive_decUsage\" bigint NOT NULL, CONSTRAINT \"UQ_fea7c0278325a1a2492f2d6acbf\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_479a8ff9d959274981087043023\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_fea7c0278325a1a2492f2d6acb\" ON \"__chart_day__instance\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__per_user_notes\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___total\" bigint NOT NULL, \"___inc\" bigint NOT NULL, \"___dec\" bigint NOT NULL, \"___diffs_normal\" bigint NOT NULL, \"___diffs_reply\" bigint NOT NULL, \"___diffs_renote\" bigint NOT NULL, CONSTRAINT \"UQ_c5545d4b31cdc684034e33b81c3\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_58bab6b6d3ad9310cbc7460fd28\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_c5545d4b31cdc684034e33b81c\" ON \"__chart_day__per_user_notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__drive\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___local_totalCount\" bigint NOT NULL, \"___local_totalSize\" bigint NOT NULL, \"___local_incCount\" bigint NOT NULL, \"___local_incSize\" bigint NOT NULL, \"___local_decCount\" bigint NOT NULL, \"___local_decSize\" bigint NOT NULL, \"___remote_totalCount\" bigint NOT NULL, \"___remote_totalSize\" bigint NOT NULL, \"___remote_incCount\" bigint NOT NULL, \"___remote_incSize\" bigint NOT NULL, \"___remote_decCount\" bigint NOT NULL, \"___remote_decSize\" bigint NOT NULL, CONSTRAINT \"UQ_0b60ebb3aa0065f10b0616c1171\" UNIQUE (\"date\"), CONSTRAINT \"PK_e7ec0de057c77c40fc8d8b62151\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0b60ebb3aa0065f10b0616c117\" ON \"__chart_day__drive\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__per_user_reaction\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___local_count\" bigint NOT NULL, \"___remote_count\" bigint NOT NULL, CONSTRAINT \"UQ_d54b653660d808b118e36c184c0\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_8af24e2d51ff781a354fe595eda\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_d54b653660d808b118e36c184c\" ON \"__chart_day__per_user_reaction\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__hashtag\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___local_users\" character varying array NOT NULL, \"___remote_users\" character varying array NOT NULL, CONSTRAINT \"UQ_8f589cf056ff51f09d6096f6450\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_13d5a3b089344e5557f8e0980b4\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_8f589cf056ff51f09d6096f645\" ON \"__chart_day__hashtag\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__per_user_following\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___local_followings_total\" bigint NOT NULL, \"___local_followings_inc\" bigint NOT NULL, \"___local_followings_dec\" bigint NOT NULL, \"___local_followers_total\" bigint NOT NULL, \"___local_followers_inc\" bigint NOT NULL, \"___local_followers_dec\" bigint NOT NULL, \"___remote_followings_total\" bigint NOT NULL, \"___remote_followings_inc\" bigint NOT NULL, \"___remote_followings_dec\" bigint NOT NULL, \"___remote_followers_total\" bigint NOT NULL, \"___remote_followers_inc\" bigint NOT NULL, \"___remote_followers_dec\" bigint NOT NULL, CONSTRAINT \"UQ_e4849a3231f38281280ea4c0eee\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_68ce6b67da57166da66fc8fb27e\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_e4849a3231f38281280ea4c0ee\" ON \"__chart_day__per_user_following\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__per_user_drive\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"group\" character varying(128) NOT NULL, \"___totalCount\" bigint NOT NULL, \"___totalSize\" bigint NOT NULL, \"___incCount\" bigint NOT NULL, \"___incSize\" bigint NOT NULL, \"___decCount\" bigint NOT NULL, \"___decSize\" bigint NOT NULL, CONSTRAINT \"UQ_62aa5047b5aec92524f24c701d7\" UNIQUE (\"date\", \"group\"), CONSTRAINT \"PK_1ae135254c137011645da7f4045\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_62aa5047b5aec92524f24c701d\" ON \"__chart_day__per_user_drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"group\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD CONSTRAINT \"UQ_36cb699c49580d4e6c2e6159f97\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ADD CONSTRAINT \"UQ_42eb716a37d381cdf566192b2be\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ADD CONSTRAINT \"UQ_845254b3eaf708ae8a6cac30265\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ADD CONSTRAINT \"UQ_a1efd3e0048a5f2793a47360dc6\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD CONSTRAINT \"UQ_0ad37b7ef50f4ddc84363d7ccca\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ALTER COLUMN \"___local_users\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ALTER COLUMN \"___remote_users\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_39ee857ab2f23493037c6b6631\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_5048e9daccbbbc6d567bb142d3\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ADD CONSTRAINT \"UQ_13565815f618a1ff53886c5b28a\" UNIQUE (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_229a41ad465f9205f1f5703291\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_25a97c02003338124b2b75fdbc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ALTER COLUMN \"___local_users\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ALTER COLUMN \"___remote_users\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_b77d4dd9562c3a899d9a286fcd\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_30bf67687f483ace115c5ca642\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"group\" SET NOT NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_36cb699c49580d4e6c2e6159f9\" ON \"__chart__federation\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_42eb716a37d381cdf566192b2b\" ON \"__chart__notes\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_845254b3eaf708ae8a6cac3026\" ON \"__chart__users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a1efd3e0048a5f2793a47360dc\" ON \"__chart__network\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0ad37b7ef50f4ddc84363d7ccc\" ON \"__chart__active_users\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_39ee857ab2f23493037c6b6631\" ON \"__chart__instance\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_5048e9daccbbbc6d567bb142d3\" ON \"__chart__per_user_notes\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_13565815f618a1ff53886c5b28\" ON \"__chart__drive\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_229a41ad465f9205f1f5703291\" ON \"__chart__per_user_reaction\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_25a97c02003338124b2b75fdbc\" ON \"__chart__hashtag\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_b77d4dd9562c3a899d9a286fcd\" ON \"__chart__per_user_following\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_30bf67687f483ace115c5ca642\" ON \"__chart__per_user_drive\" (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ADD CONSTRAINT \"UQ_39ee857ab2f23493037c6b66311\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ADD CONSTRAINT \"UQ_5048e9daccbbbc6d567bb142d34\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ADD CONSTRAINT \"UQ_229a41ad465f9205f1f57032910\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD CONSTRAINT \"UQ_25a97c02003338124b2b75fdbc8\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ADD CONSTRAINT \"UQ_b77d4dd9562c3a899d9a286fcd7\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ADD CONSTRAINT \"UQ_30bf67687f483ace115c5ca6429\" UNIQUE (\"date\", \"group\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"emoji\" RENAME COLUMN \"url\" TO \"originalUrl\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"emoji\" ADD \"publicUrl\" character varying(512) NOT NULL DEFAULT ''"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD \"webpublicType\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" DROP COLUMN \"___drive_totalUsage\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" DROP COLUMN \"___drive_totalUsage\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"___local_totalCount\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"___local_totalSize\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"___remote_totalCount\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" DROP COLUMN \"___remote_totalSize\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" DROP COLUMN \"___local_totalCount\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" DROP COLUMN \"___local_totalSize\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" DROP COLUMN \"___remote_totalCount\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" DROP COLUMN \"___remote_totalSize\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___local_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___remote_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___local_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___remote_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"___local_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"___remote_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ADD \"___local_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ADD \"___remote_users\" bigint NOT NULL DEFAULT 0"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___local_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___remote_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___local_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___remote_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"unique_temp___local_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ADD \"unique_temp___remote_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ADD \"unique_temp___local_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ADD \"unique_temp___remote_users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___incomingRequests\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___outgoingRequests\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___totalTime\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___incomingBytes\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___outgoingBytes\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___incomingRequests\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___outgoingRequests\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___totalTime\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___incomingBytes\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___outgoingBytes\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_failed\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_succeeded\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_received\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_totalFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_incFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_decFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_incUsage\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_decUsage\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_failed\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_succeeded\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_received\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_totalFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_incFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_decFiles\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_incUsage\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_decUsage\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_normal\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_reply\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_renote\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ALTER COLUMN \"___local_count\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ALTER COLUMN \"___remote_count\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_reaction\" ALTER COLUMN \"___local_count\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_reaction\" ALTER COLUMN \"___remote_count\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_total\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_inc\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_dec\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___totalCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___totalSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___totalCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___totalSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___incCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___incSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___decCount\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___decSize\" SET DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("SELECT \"userId\", \"mutedWords\" FROM \"user_profile\" WHERE \"userHost\" IS NULL"); err != nil {
				return err
			}
			if _, err := tx.Exec("/${regexp[1]}/${regexp[2]}"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__federation\" SET \"___instance_total\"=2147483647 WHERE \"___instance_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__federation\" SET \"___instance_inc\"=32767 WHERE \"___instance_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__federation\" SET \"___instance_dec\"=32767 WHERE \"___instance_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__federation\" SET \"___instance_total\"=2147483647 WHERE \"___instance_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__federation\" SET \"___instance_inc\"=32767 WHERE \"___instance_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__federation\" SET \"___instance_dec\"=32767 WHERE \"___instance_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_total\"=2147483647 WHERE \"___local_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_inc\"=2147483647 WHERE \"___local_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_dec\"=2147483647 WHERE \"___local_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_diffs_normal\"=2147483647 WHERE \"___local_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_diffs_reply\"=2147483647 WHERE \"___local_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___local_diffs_renote\"=2147483647 WHERE \"___local_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_total\"=2147483647 WHERE \"___remote_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_inc\"=2147483647 WHERE \"___remote_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_dec\"=2147483647 WHERE \"___remote_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_diffs_normal\"=2147483647 WHERE \"___remote_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_diffs_reply\"=2147483647 WHERE \"___remote_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__notes\" SET \"___remote_diffs_renote\"=2147483647 WHERE \"___remote_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_total\"=2147483647 WHERE \"___local_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_inc\"=2147483647 WHERE \"___local_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_dec\"=2147483647 WHERE \"___local_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_diffs_normal\"=2147483647 WHERE \"___local_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_diffs_reply\"=2147483647 WHERE \"___local_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___local_diffs_renote\"=2147483647 WHERE \"___local_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_total\"=2147483647 WHERE \"___remote_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_inc\"=2147483647 WHERE \"___remote_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_dec\"=2147483647 WHERE \"___remote_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_diffs_normal\"=2147483647 WHERE \"___remote_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_diffs_reply\"=2147483647 WHERE \"___remote_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__notes\" SET \"___remote_diffs_renote\"=2147483647 WHERE \"___remote_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___local_total\"=2147483647 WHERE \"___local_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___local_inc\"=32767 WHERE \"___local_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___local_dec\"=32767 WHERE \"___local_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___remote_total\"=2147483647 WHERE \"___remote_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___remote_inc\"=32767 WHERE \"___remote_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__users\" SET \"___remote_dec\"=32767 WHERE \"___remote_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___local_total\"=2147483647 WHERE \"___local_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___local_inc\"=32767 WHERE \"___local_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___local_dec\"=32767 WHERE \"___local_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___remote_total\"=2147483647 WHERE \"___remote_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___remote_inc\"=32767 WHERE \"___remote_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__users\" SET \"___remote_dec\"=32767 WHERE \"___remote_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__network\" SET \"___incomingRequests\"=2147483647 WHERE \"___incomingRequests\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__network\" SET \"___outgoingRequests\"=2147483647 WHERE \"___outgoingRequests\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__network\" SET \"___totalTime\"=2147483647 WHERE \"___totalTime\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__network\" SET \"___incomingBytes\"=2147483647 WHERE \"___incomingBytes\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__network\" SET \"___outgoingBytes\"=2147483647 WHERE \"___outgoingBytes\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__network\" SET \"___incomingRequests\"=2147483647 WHERE \"___incomingRequests\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__network\" SET \"___outgoingRequests\"=2147483647 WHERE \"___outgoingRequests\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__network\" SET \"___totalTime\"=2147483647 WHERE \"___totalTime\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__network\" SET \"___incomingBytes\"=2147483647 WHERE \"___incomingBytes\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__network\" SET \"___outgoingBytes\"=2147483647 WHERE \"___outgoingBytes\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___requests_failed\"=32767 WHERE \"___requests_failed\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___requests_succeeded\"=32767 WHERE \"___requests_succeeded\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___requests_received\"=32767 WHERE \"___requests_received\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_total\"=2147483647 WHERE \"___notes_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_inc\"=2147483647 WHERE \"___notes_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_dec\"=2147483647 WHERE \"___notes_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_diffs_normal\"=2147483647 WHERE \"___notes_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_diffs_reply\"=2147483647 WHERE \"___notes_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___notes_diffs_renote\"=2147483647 WHERE \"___notes_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___users_total\"=2147483647 WHERE \"___users_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___users_inc\"=32767 WHERE \"___users_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___users_dec\"=32767 WHERE \"___users_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___following_total\"=2147483647 WHERE \"___following_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___following_inc\"=32767 WHERE \"___following_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___following_dec\"=32767 WHERE \"___following_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___followers_total\"=2147483647 WHERE \"___followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___followers_inc\"=32767 WHERE \"___followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___followers_dec\"=32767 WHERE \"___followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___drive_totalFiles\"=2147483647 WHERE \"___drive_totalFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___drive_incFiles\"=2147483647 WHERE \"___drive_incFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___drive_decFiles\"=2147483647 WHERE \"___drive_decFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___drive_incUsage\"=2147483647 WHERE \"___drive_incUsage\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__instance\" SET \"___drive_decUsage\"=2147483647 WHERE \"___drive_decUsage\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___requests_failed\"=32767 WHERE \"___requests_failed\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___requests_succeeded\"=32767 WHERE \"___requests_succeeded\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___requests_received\"=32767 WHERE \"___requests_received\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_total\"=2147483647 WHERE \"___notes_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_inc\"=2147483647 WHERE \"___notes_inc\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_dec\"=2147483647 WHERE \"___notes_dec\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_diffs_normal\"=2147483647 WHERE \"___notes_diffs_normal\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_diffs_reply\"=2147483647 WHERE \"___notes_diffs_reply\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___notes_diffs_renote\"=2147483647 WHERE \"___notes_diffs_renote\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___users_total\"=2147483647 WHERE \"___users_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___users_inc\"=32767 WHERE \"___users_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___users_dec\"=32767 WHERE \"___users_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___following_total\"=2147483647 WHERE \"___following_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___following_inc\"=32767 WHERE \"___following_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___following_dec\"=32767 WHERE \"___following_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___followers_total\"=2147483647 WHERE \"___followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___followers_inc\"=32767 WHERE \"___followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___followers_dec\"=32767 WHERE \"___followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___drive_totalFiles\"=2147483647 WHERE \"___drive_totalFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___drive_incFiles\"=2147483647 WHERE \"___drive_incFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___drive_decFiles\"=2147483647 WHERE \"___drive_decFiles\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___drive_incUsage\"=2147483647 WHERE \"___drive_incUsage\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__instance\" SET \"___drive_decUsage\"=2147483647 WHERE \"___drive_decUsage\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___total\"=2147483647 WHERE \"___total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___inc\"=32767 WHERE \"___inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___dec\"=32767 WHERE \"___dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___diffs_normal\"=32767 WHERE \"___diffs_normal\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___diffs_reply\"=32767 WHERE \"___diffs_reply\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_notes\" SET \"___diffs_renote\"=32767 WHERE \"___diffs_renote\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___total\"=2147483647 WHERE \"___total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___inc\"=32767 WHERE \"___inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___dec\"=32767 WHERE \"___dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___diffs_normal\"=32767 WHERE \"___diffs_normal\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___diffs_reply\"=32767 WHERE \"___diffs_reply\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_notes\" SET \"___diffs_renote\"=32767 WHERE \"___diffs_renote\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___local_incCount\"=2147483647 WHERE \"___local_incCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___local_incSize\"=2147483647 WHERE \"___local_incSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___local_decCount\"=2147483647 WHERE \"___local_decCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___local_decSize\"=2147483647 WHERE \"___local_decSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___remote_incCount\"=2147483647 WHERE \"___remote_incCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___remote_incSize\"=2147483647 WHERE \"___remote_incSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___remote_decCount\"=2147483647 WHERE \"___remote_decCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__drive\" SET \"___remote_decSize\"=2147483647 WHERE \"___remote_decSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___local_incCount\"=2147483647 WHERE \"___local_incCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___local_incSize\"=2147483647 WHERE \"___local_incSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___local_decCount\"=2147483647 WHERE \"___local_decCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___local_decSize\"=2147483647 WHERE \"___local_decSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___remote_incCount\"=2147483647 WHERE \"___remote_incCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___remote_incSize\"=2147483647 WHERE \"___remote_incSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___remote_decCount\"=2147483647 WHERE \"___remote_decCount\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__drive\" SET \"___remote_decSize\"=2147483647 WHERE \"___remote_decSize\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_reaction\" SET \"___local_count\"=32767 WHERE \"___local_count\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_reaction\" SET \"___remote_count\"=32767 WHERE \"___remote_count\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_reaction\" SET \"___local_count\"=32767 WHERE \"___local_count\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_reaction\" SET \"___remote_count\"=32767 WHERE \"___remote_count\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followings_total\"=2147483647 WHERE \"___local_followings_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followings_inc\"=32767 WHERE \"___local_followings_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followings_dec\"=32767 WHERE \"___local_followings_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followers_total\"=2147483647 WHERE \"___local_followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followers_inc\"=32767 WHERE \"___local_followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___local_followers_dec\"=32767 WHERE \"___local_followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followings_total\"=2147483647 WHERE \"___remote_followings_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followings_inc\"=32767 WHERE \"___remote_followings_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followings_dec\"=32767 WHERE \"___remote_followings_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followers_total\"=2147483647 WHERE \"___remote_followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followers_inc\"=32767 WHERE \"___remote_followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__per_user_following\" SET \"___remote_followers_dec\"=32767 WHERE \"___remote_followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followings_total\"=2147483647 WHERE \"___local_followings_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followings_inc\"=32767 WHERE \"___local_followings_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followings_dec\"=32767 WHERE \"___local_followings_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followers_total\"=2147483647 WHERE \"___local_followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followers_inc\"=32767 WHERE \"___local_followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___local_followers_dec\"=32767 WHERE \"___local_followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followings_total\"=2147483647 WHERE \"___remote_followings_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followings_inc\"=32767 WHERE \"___remote_followings_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followings_dec\"=32767 WHERE \"___remote_followings_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followers_total\"=2147483647 WHERE \"___remote_followers_total\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followers_inc\"=32767 WHERE \"___remote_followers_inc\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__per_user_following\" SET \"___remote_followers_dec\"=32767 WHERE \"___remote_followers_dec\" > 32767"); err != nil {
				return err
			}
			if _, err := tx.Exec("TRUNCATE TABLE \"__chart__per_user_drive\""); err != nil {
				return err
			}
			if _, err := tx.Exec("TRUNCATE TABLE \"__chart_day__per_user_drive\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_total\" TYPE integer USING \"___instance_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_inc\" TYPE smallint USING \"___instance_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ALTER COLUMN \"___instance_dec\" TYPE smallint USING \"___instance_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_total\" TYPE integer USING \"___instance_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_inc\" TYPE smallint USING \"___instance_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ALTER COLUMN \"___instance_dec\" TYPE smallint USING \"___instance_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_total\" TYPE integer USING \"___local_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_inc\" TYPE integer USING \"___local_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_dec\" TYPE integer USING \"___local_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_normal\" TYPE integer USING \"___local_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_reply\" TYPE integer USING \"___local_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___local_diffs_renote\" TYPE integer USING \"___local_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_total\" TYPE integer USING \"___remote_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_inc\" TYPE integer USING \"___remote_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_dec\" TYPE integer USING \"___remote_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_normal\" TYPE integer USING \"___remote_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_reply\" TYPE integer USING \"___remote_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ALTER COLUMN \"___remote_diffs_renote\" TYPE integer USING \"___remote_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_total\" TYPE integer USING \"___local_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_inc\" TYPE integer USING \"___local_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_dec\" TYPE integer USING \"___local_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_normal\" TYPE integer USING \"___local_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_reply\" TYPE integer USING \"___local_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___local_diffs_renote\" TYPE integer USING \"___local_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_total\" TYPE integer USING \"___remote_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_inc\" TYPE integer USING \"___remote_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_dec\" TYPE integer USING \"___remote_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_normal\" TYPE integer USING \"___remote_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_reply\" TYPE integer USING \"___remote_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ALTER COLUMN \"___remote_diffs_renote\" TYPE integer USING \"___remote_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_total\" TYPE integer USING \"___local_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_inc\" TYPE smallint USING \"___local_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___local_dec\" TYPE smallint USING \"___local_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_total\" TYPE integer USING \"___remote_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_inc\" TYPE smallint USING \"___remote_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__users\" ALTER COLUMN \"___remote_dec\" TYPE smallint USING \"___remote_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_total\" TYPE integer USING \"___local_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_inc\" TYPE smallint USING \"___local_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___local_dec\" TYPE smallint USING \"___local_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_total\" TYPE integer USING \"___remote_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_inc\" TYPE smallint USING \"___remote_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__users\" ALTER COLUMN \"___remote_dec\" TYPE smallint USING \"___remote_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___incomingRequests\" TYPE integer USING \"___incomingRequests\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___outgoingRequests\" TYPE integer USING \"___outgoingRequests\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___totalTime\" TYPE integer USING \"___totalTime\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___incomingBytes\" TYPE integer USING \"___incomingBytes\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__network\" ALTER COLUMN \"___outgoingBytes\" TYPE integer USING \"___outgoingBytes\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___incomingRequests\" TYPE integer USING \"___incomingRequests\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___outgoingRequests\" TYPE integer USING \"___outgoingRequests\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___totalTime\" TYPE integer USING \"___totalTime\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___incomingBytes\" TYPE integer USING \"___incomingBytes\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__network\" ALTER COLUMN \"___outgoingBytes\" TYPE integer USING \"___outgoingBytes\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_failed\" TYPE smallint USING \"___requests_failed\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_succeeded\" TYPE smallint USING \"___requests_succeeded\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___requests_received\" TYPE smallint USING \"___requests_received\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_total\" TYPE integer USING \"___notes_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_inc\" TYPE integer USING \"___notes_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_dec\" TYPE integer USING \"___notes_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_normal\" TYPE integer USING \"___notes_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_reply\" TYPE integer USING \"___notes_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___notes_diffs_renote\" TYPE integer USING \"___notes_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_total\" TYPE integer USING \"___users_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_inc\" TYPE smallint USING \"___users_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___users_dec\" TYPE smallint USING \"___users_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_total\" TYPE integer USING \"___following_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_inc\" TYPE smallint USING \"___following_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___following_dec\" TYPE smallint USING \"___following_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_total\" TYPE integer USING \"___followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_inc\" TYPE smallint USING \"___followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___followers_dec\" TYPE smallint USING \"___followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_totalFiles\" TYPE integer USING \"___drive_totalFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_incFiles\" TYPE integer USING \"___drive_incFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_decFiles\" TYPE integer USING \"___drive_decFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_incUsage\" TYPE integer USING \"___drive_incUsage\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ALTER COLUMN \"___drive_decUsage\" TYPE integer USING \"___drive_decUsage\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_failed\" TYPE smallint USING \"___requests_failed\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_succeeded\" TYPE smallint USING \"___requests_succeeded\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___requests_received\" TYPE smallint USING \"___requests_received\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_total\" TYPE integer USING \"___notes_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_inc\" TYPE integer USING \"___notes_inc\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_dec\" TYPE integer USING \"___notes_dec\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_normal\" TYPE integer USING \"___notes_diffs_normal\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_reply\" TYPE integer USING \"___notes_diffs_reply\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___notes_diffs_renote\" TYPE integer USING \"___notes_diffs_renote\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_total\" TYPE integer USING \"___users_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_inc\" TYPE smallint USING \"___users_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___users_dec\" TYPE smallint USING \"___users_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_total\" TYPE integer USING \"___following_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_inc\" TYPE smallint USING \"___following_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___following_dec\" TYPE smallint USING \"___following_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_total\" TYPE integer USING \"___followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_inc\" TYPE smallint USING \"___followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___followers_dec\" TYPE smallint USING \"___followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_totalFiles\" TYPE integer USING \"___drive_totalFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_incFiles\" TYPE integer USING \"___drive_incFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_decFiles\" TYPE integer USING \"___drive_decFiles\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_incUsage\" TYPE integer USING \"___drive_incUsage\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ALTER COLUMN \"___drive_decUsage\" TYPE integer USING \"___drive_decUsage\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___total\" TYPE integer USING \"___total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___inc\" TYPE smallint USING \"___inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___dec\" TYPE smallint USING \"___dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_normal\" TYPE smallint USING \"___diffs_normal\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_reply\" TYPE smallint USING \"___diffs_reply\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ALTER COLUMN \"___diffs_renote\" TYPE smallint USING \"___diffs_renote\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___total\" TYPE integer USING \"___total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___inc\" TYPE smallint USING \"___inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___dec\" TYPE smallint USING \"___dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_normal\" TYPE smallint USING \"___diffs_normal\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_reply\" TYPE smallint USING \"___diffs_reply\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ALTER COLUMN \"___diffs_renote\" TYPE smallint USING \"___diffs_renote\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_incCount\" TYPE integer USING \"___local_incCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_incSize\" TYPE integer USING \"___local_incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_decCount\" TYPE integer USING \"___local_decCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___local_decSize\" TYPE integer USING \"___local_decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_incCount\" TYPE integer USING \"___remote_incCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_incSize\" TYPE integer USING \"___remote_incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_decCount\" TYPE integer USING \"___remote_decCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__drive\" ALTER COLUMN \"___remote_decSize\" TYPE integer USING \"___remote_decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_incCount\" TYPE integer USING \"___local_incCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_incSize\" TYPE integer USING \"___local_incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_decCount\" TYPE integer USING \"___local_decCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___local_decSize\" TYPE integer USING \"___local_decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_incCount\" TYPE integer USING \"___remote_incCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_incSize\" TYPE integer USING \"___remote_incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_decCount\" TYPE integer USING \"___remote_decCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__drive\" ALTER COLUMN \"___remote_decSize\" TYPE integer USING \"___remote_decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ALTER COLUMN \"___local_count\" TYPE smallint USING \"___local_count\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_reaction\" ALTER COLUMN \"___remote_count\" TYPE smallint USING \"___remote_count\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_reaction\" ALTER COLUMN \"___local_count\" TYPE smallint USING \"___local_count\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_reaction\" ALTER COLUMN \"___remote_count\" TYPE smallint USING \"___remote_count\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_total\" TYPE integer USING \"___local_followings_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_inc\" TYPE smallint USING \"___local_followings_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followings_dec\" TYPE smallint USING \"___local_followings_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_total\" TYPE integer USING \"___local_followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_inc\" TYPE smallint USING \"___local_followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___local_followers_dec\" TYPE smallint USING \"___local_followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_total\" TYPE integer USING \"___remote_followings_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_inc\" TYPE smallint USING \"___remote_followings_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followings_dec\" TYPE smallint USING \"___remote_followings_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_total\" TYPE integer USING \"___remote_followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_inc\" TYPE smallint USING \"___remote_followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_following\" ALTER COLUMN \"___remote_followers_dec\" TYPE smallint USING \"___remote_followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_total\" TYPE integer USING \"___local_followings_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_inc\" TYPE smallint USING \"___local_followings_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followings_dec\" TYPE smallint USING \"___local_followings_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_total\" TYPE integer USING \"___local_followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_inc\" TYPE smallint USING \"___local_followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___local_followers_dec\" TYPE smallint USING \"___local_followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_total\" TYPE integer USING \"___remote_followings_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_inc\" TYPE smallint USING \"___remote_followings_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followings_dec\" TYPE smallint USING \"___remote_followings_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_total\" TYPE integer USING \"___remote_followers_total\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_inc\" TYPE smallint USING \"___remote_followers_inc\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_following\" ALTER COLUMN \"___remote_followers_dec\" TYPE smallint USING \"___remote_followers_dec\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___totalCount\" TYPE integer USING \"___totalCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___totalSize\" TYPE integer USING \"___totalSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___incCount\" TYPE smallint USING \"___incCount\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___incSize\" TYPE integer USING \"___incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___decCount\" TYPE smallint USING \"___decCount\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_drive\" ALTER COLUMN \"___decSize\" TYPE integer USING \"___decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___totalCount\" TYPE integer USING \"___totalCount\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___totalSize\" TYPE integer USING \"___totalSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___incCount\" TYPE smallint USING \"___incCount\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___incSize\" TYPE integer USING \"___incSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___decCount\" TYPE smallint USING \"___decCount\"::smallint"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_drive\" ALTER COLUMN \"___decSize\" TYPE integer USING \"___decSize\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__active_users\" SET \"___local_users\"=2147483647 WHERE \"___local_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__active_users\" SET \"___remote_users\"=2147483647 WHERE \"___remote_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__active_users\" SET \"___local_users\"=2147483647 WHERE \"___local_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__active_users\" SET \"___remote_users\"=2147483647 WHERE \"___remote_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ALTER COLUMN \"___local_users\" TYPE integer USING \"___local_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ALTER COLUMN \"___remote_users\" TYPE integer USING \"___remote_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ALTER COLUMN \"___local_users\" TYPE integer USING \"___local_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ALTER COLUMN \"___remote_users\" TYPE integer USING \"___remote_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__hashtag\" SET \"___local_users\"=2147483647 WHERE \"___local_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart__hashtag\" SET \"___remote_users\"=2147483647 WHERE \"___remote_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__hashtag\" SET \"___local_users\"=2147483647 WHERE \"___local_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("UPDATE \"__chart_day__hashtag\" SET \"___remote_users\"=2147483647 WHERE \"___remote_users\" > 2147483647"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ALTER COLUMN \"___local_users\" TYPE integer USING \"___local_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__hashtag\" ALTER COLUMN \"___remote_users\" TYPE integer USING \"___remote_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ALTER COLUMN \"___local_users\" TYPE integer USING \"___local_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__hashtag\" ALTER COLUMN \"___remote_users\" TYPE integer USING \"___remote_users\"::integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart__ap_request\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___deliverFailed\" integer NOT NULL DEFAULT '0', \"___deliverSucceeded\" integer NOT NULL DEFAULT '0', \"___inboxReceived\" integer NOT NULL DEFAULT '0', CONSTRAINT \"UQ_e56f4beac5746d44bc3e19c80d0\" UNIQUE (\"date\"), CONSTRAINT \"PK_56a25cd447c7ee08876b3baf8d8\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_e56f4beac5746d44bc3e19c80d\" ON \"__chart__ap_request\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"__chart_day__ap_request\" (\"id\" SERIAL NOT NULL, \"date\" integer NOT NULL, \"___deliverFailed\" integer NOT NULL DEFAULT '0', \"___deliverSucceeded\" integer NOT NULL DEFAULT '0', \"___inboxReceived\" integer NOT NULL DEFAULT '0', CONSTRAINT \"UQ_a848f66d6cec11980a5dd595822\" UNIQUE (\"date\"), CONSTRAINT \"PK_9318b49daee320194e23f712e69\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_a848f66d6cec11980a5dd59582\" ON \"__chart_day__ap_request\" (\"date\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"unique_temp___deliveredInstances\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___deliveredInstances\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"unique_temp___inboxInstances\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___inboxInstances\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"unique_temp___deliveredInstances\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___deliveredInstances\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"unique_temp___inboxInstances\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___inboxInstances\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"unique_temp___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"unique_temp___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"unique_temp___local_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"unique_temp___remote_users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___users\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___notedUsers\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___notedUsers\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredWithinWeek\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredWithinWeek\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredWithinMonth\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredWithinMonth\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredWithinYear\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredWithinYear\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredOutsideWeek\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredOutsideWeek\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredOutsideMonth\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredOutsideMonth\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___registeredOutsideYear\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___registeredOutsideYear\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___users\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___users\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___notedUsers\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___notedUsers\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredWithinWeek\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredWithinWeek\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredWithinMonth\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredWithinMonth\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredWithinYear\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredWithinYear\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredOutsideWeek\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredOutsideWeek\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredOutsideMonth\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredOutsideMonth\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___registeredOutsideYear\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___registeredOutsideYear\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ADD \"___local_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__notes\" ADD \"___remote_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ADD \"___local_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__notes\" ADD \"___remote_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__instance\" ADD \"___notes_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__instance\" ADD \"___notes_diffs_withFile\" integer NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__per_user_notes\" ADD \"___diffs_withFile\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__per_user_notes\" ADD \"___diffs_withFile\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"unique_temp___stalled\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___stalled\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"unique_temp___stalled\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___stalled\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"unique_temp___users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"unique_temp___notedUsers\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" DROP COLUMN \"___notedUsers\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"unique_temp___users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___users\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"unique_temp___notedUsers\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" DROP COLUMN \"___notedUsers\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___readWrite\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___read\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___read\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"unique_temp___write\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__active_users\" ADD \"___write\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___readWrite\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___read\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___read\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"unique_temp___write\" character varying array NOT NULL DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__active_users\" ADD \"___write\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"themeColor\" character varying(512)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"___instance_total\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"___instance_inc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"___instance_dec\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" DROP COLUMN \"___instance_total\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" DROP COLUMN \"___instance_inc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" DROP COLUMN \"___instance_dec\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___sub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___pub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___sub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___pub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_4ccd2239268ebbd1b35e318754\" ON \"following\" (\"followerHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_fcdafee716dfe9c3b5fde90f30\" ON \"following\" (\"followeeHost\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" DROP COLUMN \"maxNoteTextLength\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___pubsub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___pubsub\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"defaultLightTheme\" character varying(8192)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"defaultDarkTheme\" character varying(8192)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"muting\" ADD \"expiresAt\" TIMESTAMP WITH TIME ZONE"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_c1fd1c3dfb0627aa36c253fd14\" ON \"muting\" (\"expiresAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"notification_type_enum\" RENAME TO \"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TYPE \"public\".\"notification_type_enum\" AS ENUM('follow', 'mention', 'reply', 'renote', 'quote', 'reaction', 'pollVote', 'pollEnded', 'receiveFollowRequest', 'followRequestAccepted', 'groupInvited', 'app')"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"notification\" ALTER COLUMN \"type\" TYPE \"public\".\"notification_type_enum\" USING \"type\"::\"text\"::\"public\".\"notification_type_enum\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP TYPE \"public\".\"notification_type_enum_old\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___active\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___active\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" DROP COLUMN \"driveUsage\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" DROP COLUMN \"driveFiles\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" DROP COLUMN \"___active\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" DROP COLUMN \"___active\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___subActive\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart__federation\" ADD \"___pubActive\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___subActive\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"__chart_day__federation\" ADD \"___pubActive\" smallint NOT NULL DEFAULT '0'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"webhook\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"name\" character varying(128) NOT NULL, \"on\" character varying(128) array NOT NULL DEFAULT '{}', \"url\" character varying(1024) NOT NULL, \"secret\" character varying(1024) NOT NULL, \"active\" boolean NOT NULL DEFAULT true, CONSTRAINT \"PK_e6765510c2d078db49632b59020\" PRIMARY KEY (\"id\")); COMMENT ON COLUMN \"webhook\".\"createdAt\" IS 'The created date of the Antenna.'; COMMENT ON COLUMN \"webhook\".\"userId\" IS 'The owner ID.'; COMMENT ON COLUMN \"webhook\".\"name\" IS 'The name of the Antenna.'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f272c8c8805969e6a6449c77b3\" ON \"webhook\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_8063a0586ed1dfbe86e982d961\" ON \"webhook\" (\"on\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_5a056076f76b2efe08216ba655\" ON \"webhook\" (\"active\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"webhook\" ADD CONSTRAINT \"FK_f272c8c8805969e6a6449c77b3c\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"webhook\" ADD \"latestSentAt\" TIMESTAMP WITH TIME ZONE"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"webhook\" ADD \"latestStatus\" integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER INDEX \"public\".\"IDX_seoignmeoprigmkpodgrjmkpormg\" RENAME TO \"IDX_c8cc87bd0f2f4487d17c651fbf\""); err != nil {
				return err
			}
			if _, err := tx.Exec("DROP INDEX \"public\".\"IDX_note_on_channelId_and_id_desc\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ALTER COLUMN \"followersUri\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"session\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"appId\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"name\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"description\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"access_token\" ALTER COLUMN \"iconUrl\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"softwareName\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"softwareVersion\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"name\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"description\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"maintainerName\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"maintainerEmail\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"iconUrl\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"faviconUrl\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"instance\" ALTER COLUMN \"themeColor\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"clip\" ALTER COLUMN \"description\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ALTER COLUMN \"channelId\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ALTER COLUMN \"comment\" DROP DEFAULT"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_315c779174fe8247ab324f036e\" ON \"drive_file\" (\"isLink\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_f22169eb10657bded6d875ac8f\" ON \"note\" (\"channelId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_a9021cc2e1feb5f72d3db6e9f5\" ON \"abuse_user_report\" (\"targetUserId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"abuse_user_report\" WHERE \"targetUserId\" NOT IN (SELECT \"id\" FROM \"user\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD CONSTRAINT \"FK_a9021cc2e1feb5f72d3db6e9f5f\" FOREIGN KEY (\"targetUserId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"poll\" ADD CONSTRAINT \"UQ_da851e06d0dfe2ef397d8b1bf1b\" UNIQUE (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_keypair\" ADD CONSTRAINT \"UQ_f4853eb41ab722fe05f81cedeb6\" UNIQUE (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD CONSTRAINT \"UQ_51cb79b5555effaf7d69ba1cff9\" UNIQUE (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_publickey\" ADD CONSTRAINT \"UQ_10c146e4b39b443ede016f6736d\" UNIQUE (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"promo_note\" ADD CONSTRAINT \"UQ_e263909ca4fe5d57f8d4230dd5c\" UNIQUE (\"noteId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"page\" RENAME CONSTRAINT \"FK_3126dd7c502c9e4d7597ef7ef10\" TO \"FK_a9ca79ad939bf06066b81c9d3aa\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TYPE \"public\".\"user_profile_mutingnotificationtypes_enum\" ADD VALUE 'pollEnded' AFTER 'pollVote'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user\" ADD \"driveCapacityOverrideMb\" integer"); err != nil {
				return err
			}
			if _, err := tx.Exec("COMMENT ON COLUMN \"user\".\"driveCapacityOverrideMb\" IS 'Overrides user drive capacity limit'"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"user_ip\" (\"id\" SERIAL NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"userId\" character varying(32) NOT NULL, \"ip\" character varying(128) NOT NULL, CONSTRAINT \"PK_2c44ddfbf7c0464d028dcef325e\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_7f7f1c66f48e9a8e18a33bc515\" ON \"user_ip\" (\"userId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_361b500e06721013c124b7b6c5\" ON \"user_ip\" (\"userId\", \"ip\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_ip\" ADD CONSTRAINT \"FK_7f7f1c66f48e9a8e18a33bc5150\" FOREIGN KEY (\"userId\") REFERENCES \"user\"(\"id\") ON DELETE NO ACTION ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD \"requestHeaders\" jsonb DEFAULT '{}'"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"drive_file\" ADD \"requestIp\" character varying(128)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_ip\" DROP CONSTRAINT \"FK_7f7f1c66f48e9a8e18a33bc5150\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"enableIpLogging\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"moderationNote\" character varying(8192) NOT NULL DEFAULT ''"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"enableActiveEmailValidation\" boolean NOT NULL DEFAULT true"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX IF NOT EXISTS \"IDX_a9021cc2e1feb5f72d3db6e9f5\" ON \"abuse_user_report\" (\"targetUserId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"abuse_user_report\" WHERE \"targetUserId\" NOT IN (SELECT \"id\" FROM \"user\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" DROP CONSTRAINT IF EXISTS \"FK_a9021cc2e1feb5f72d3db6e9f5f\""); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"abuse_user_report\" ADD CONSTRAINT \"FK_a9021cc2e1feb5f72d3db6e9f5f\" FOREIGN KEY (\"targetUserId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"blockedEmailDomains\" character varying(256) array NOT NULL DEFAULT '{}'::varchar[]"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"user_profile\" ADD \"twoFactorBackupSecret\" character varying array"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE TABLE \"renote_muting\" (\"id\" character varying(32) NOT NULL, \"createdAt\" TIMESTAMP WITH TIME ZONE NOT NULL, \"muteeId\" character varying(32) NOT NULL, \"muterId\" character varying(32) NOT NULL, CONSTRAINT \"PK_renoteMuting_id\" PRIMARY KEY (\"id\"))"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_renote_muting_createdAt\" ON \"muting\" (\"createdAt\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_renote_muting_muteeId\" ON \"muting\" (\"muteeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE INDEX \"IDX_renote_muting_muterId\" ON \"muting\" (\"muterId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"enableTurnstile\" boolean NOT NULL DEFAULT false"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"turnstileSiteKey\" character varying(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"meta\" ADD \"turnstileSecretKey\" character varying(64)"); err != nil {
				return err
			}
			if _, err := tx.Exec("CREATE UNIQUE INDEX \"IDX_0d801c609cec4e9eb4b6b4490c\" ON \"renote_muting\" (\"muterId\", \"muteeId\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"renote_muting\" WHERE \"muteeId\" NOT IN (SELECT \"id\" FROM \"user\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("DELETE FROM \"renote_muting\" WHERE \"muterId\" NOT IN (SELECT \"id\" FROM \"user\")"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"renote_muting\" ADD CONSTRAINT \"FK_7eac97594bcac5ffcf2068089b6\" FOREIGN KEY (\"muteeId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"renote_muting\" ADD CONSTRAINT \"FK_7aa72a5fe76019bfe8e5e0e8b7d\" FOREIGN KEY (\"muterId\") REFERENCES \"user\"(\"id\") ON DELETE CASCADE ON UPDATE NO ACTION"); err != nil {
				return err
			}
			if _, err := tx.Exec("ALTER TABLE \"note\" ADD \"updatedAt\" TIMESTAMP WITH TIME ZONE"); err != nil {
				return err
			}

			return nil
		}})
