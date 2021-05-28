package restclient

// Discord Endpoint Constants
const (
	APIVersion = "9"
	Base       = "https://discord.com/"
	CDN        = "https://cdn.discordapp.com"
	API        = Base + "api/v" + APIVersion
)

// Misc
var (
	GetGateway        = NewAPIRoute(GET, "/gateway")
	GetGatewayBot     = NewAPIRoute(GET, "/gateway/bot")
	GetBotApplication = NewAPIRoute(GET, "/oauth2/applications/@me")
	GetVoiceRegions   = NewAPIRoute(GET, "/voice/regions")
)

// Users
var (
	GetUser         = NewAPIRoute(GET, "/users/{user.id}")
	GetSelfUser     = NewAPIRoute(GET, "/users/@me")
	UpdateSelfUser  = NewAPIRoute(PATCH, "/users/@me")
	GetGuilds       = NewAPIRoute(GET, "/users/@me/guilds")
	LeaveGuild      = NewAPIRoute(DELETE, "/users/@me/guilds/{guild.id}")
	GetDMChannels   = NewAPIRoute(GET, "/users/@me/channels")
	CreateDMChannel = NewAPIRoute(POST, "/users/@me/channels")
)

// Guilds
var (
	GetGuild          = NewAPIRoute(GET, "/guilds/{guild.id}")
	CreateGuild       = NewAPIRoute(POST, "/guilds")
	UpdateGuild       = NewAPIRoute(PATCH, "/guilds/{guild.id}")
	DeleteGuild       = NewAPIRoute(DELETE, "/guilds/{guild.id}")
	GetGuildVanityURL = NewAPIRoute(GET, "/guilds/{guild.id}/vanity-url")

	CreateGuildChannel     = NewAPIRoute(POST, "/guilds/{guild.id}/channels")
	GetGuildChannels       = NewAPIRoute(GET, "/guilds/{guild.id}/channels")
	UpdateChannelPositions = NewAPIRoute(PATCH, "/guilds/{guild.id}/channels")

	GetBans   = NewAPIRoute(GET, "/guilds/{guild.id}/bans")
	GetBan    = NewAPIRoute(GET, "/guilds/{guild.id}/bans/{user.id}")
	CreateBan = NewAPIRoute(POST, "/guilds/{guild.id}/bans/{user.id}", "delete_message_days", "reason")
	DeleteBan = NewAPIRoute(DELETE, "/guilds/{guild.id}/bans/{user.id}")

	GetMember        = NewAPIRoute(GET, "/guilds/{guild.id}/members/{user.id}")
	GetMembers       = NewAPIRoute(GET, "/guilds/{guild.id}/members")
	AddMember        = NewAPIRoute(PUT, "/guilds/{guild.id}/members/{user.id}")
	UpdateMember     = NewAPIRoute(PATCH, "/guilds/{guild.id}/members/{user.id}")
	RemoveMember     = NewAPIRoute(DELETE, "/guilds/{guild.id}/members/{user.id}", "reason")
	AddMemberRole    = NewAPIRoute(PUT, "/guilds/{guild.id}/members/{user.id}/roles/{role.id}")
	RemoveMemberRole = NewAPIRoute(DELETE, "/guilds/{guild.id}/members/{user.id}/roles/{role.id}")

	UpdateSelfNick = NewAPIRoute(PATCH, "/guilds/{guild.id}/members/@me/nick")

	GetPruneMembersCount = NewAPIRoute(GET, "/guilds/{guild.id}/prune")
	PruneMembers         = NewAPIRoute(POST, "/guilds/{guild.id}/prune")

	GetGuildWebhooks = NewAPIRoute(GET, "/guilds/{guild.id}/webhooks")

	GetGuildAuditLogs = NewAPIRoute(GET, "/guilds/{guild.id}/audit-logs")

	GetGuildVoiceRegions = NewAPIRoute(GET, "/guilds/{guild.id}/regions")

	GetGuildIntegrations   = NewAPIRoute(GET, "/guilds/{guild.id}/integrations")
	CreateGuildIntegration = NewAPIRoute(POST, "/guilds/{guild.id}/integrations")
	UpdateGuildIntegration = NewAPIRoute(PATCH, "/guilds/{guild.id}/integrations/{integration.id}")
	DeleteGuildIntegration = NewAPIRoute(DELETE, "/guilds/{guild.id}/integrations/{integration.id}")
	SyncGuildIntegration   = NewAPIRoute(POST, "/guilds/{guild.id}/integrations/{integration.id}/sync")
)

// Roles
var (
	GetRoles            = NewAPIRoute(GET, "/guilds/{guild.id}/roles")
	GetRole             = NewAPIRoute(GET, "/guilds/{guild.id}/roles/{role.id}")
	CreateRole          = NewAPIRoute(POST, "/guilds/{guild.id}/roles")
	UpdateRoles         = NewAPIRoute(PATCH, "/guilds/{guild.id}/roles")
	UpdateRole          = NewAPIRoute(PATCH, "/guilds/{guild.id}/roles/{role.id}")
	UpdateRolePositions = NewAPIRoute(PATCH, "/guilds/{guild.id}/roles")
	DelteRole           = NewAPIRoute(DELETE, "/guilds/{guild.id}/roles/{role.id}")
)

// Channels
var (
	GetChannel    = NewAPIRoute(GET, "/channels/{channel.id}")
	UpdateChannel = NewAPIRoute(PATCH, "/channels/{channel.id}")
	DeleteChannel = NewAPIRoute(DELETE, "/channels/{channel.id}")

	GetWebhooks   = NewAPIRoute(GET, "/channels/{channel.id}/webhooks")
	CreateWebhook = NewAPIRoute(POST, "/channels/{channel.id}/webhooks")

	GetPermissionOverrides   = NewAPIRoute(GET, "/channels/{channel.id}/permissions")
	GetPermissionOverride    = NewAPIRoute(GET, "/channels/{channel.id}/permissions/{overwrite.id}")
	CreatePermissionOverride = NewAPIRoute(PUT, "/channels/{channel.id}/permissions/{overwrite.id}")
	UpdatePermissionOverride = NewAPIRoute(PUT, "/channels/{channel.id}/permissions/{overwrite.id}")
	DeletePermissionOverride = NewAPIRoute(DELETE, "/channels/{channel.id}/permissions/{overwrite.id}")

	SendTyping = NewAPIRoute(POST, "/channels/{channel.id}/typing")
)

// Threads
var (
	CreateThreadWithMessage = NewAPIRoute(POST, "/channels/{channel.id}/messages/{message.id}/threads")
	CreateThread            = NewAPIRoute(POST, "/channels/{channel.id}/threads")
	JoinThread              = NewAPIRoute(PUT, "/channels/{channel.id}/thread-members/@me")
	LeaveThread             = NewAPIRoute(DELETE, "/channels/{channel.id}/thread-members/@me")
	AddThreadMember         = NewAPIRoute(PUT, "/channels/{channel.id}/thread-members/{user.id}")
	RemoveThreadMember      = NewAPIRoute(DELETE, "/channels/{channel.id}/thread-members/{user.id}")
	GetThreadMembers        = NewAPIRoute(GET, "/channels/{channel.id}/thread-members")

	GetActiveThreads                 = NewAPIRoute(GET, "/channels/{channel.id}/threads/active")
	GetArchivedPublicThreads         = NewAPIRoute(GET, "/channels/{channel.id}/threads/archived/public")
	GetArchivedPrivateThreads        = NewAPIRoute(GET, "/channels/{channel.id}/threads/archived/private")
	GetJoinedArchievedPrivateThreads = NewAPIRoute(GET, "/channels/{channel.id}/users/@me/threads/archived/private")
)

// Messages
var (
	GetMessages       = NewAPIRoute(GET, "/channels/{channel.id}/messages")
	GetMessage        = NewAPIRoute(GET, "/channels/{channel.id}/messages/{message.id}")
	CreateMessage     = NewAPIRoute(POST, "/channels/{channel.id}/messages")
	UpdateMessage     = NewAPIRoute(PATCH, "/channels/{channel.id}/messages/{message.id}")
	DeleteMessage     = NewAPIRoute(DELETE, "/channels/{channel.id}/messages/{message.id}")
	BulkDeleteMessage = NewAPIRoute(POST, "/channels/{channel.id}/messages/bulk-delete")

	GetPinnedMessages   = NewAPIRoute(GET, "/channels/{channel.id}/pins")
	AddPinnedMessage    = NewAPIRoute(PUT, "/channels/{channel.id}/pins/{message.id}")
	RemovePinnedMessage = NewAPIRoute(DELETE, "/channels/{channel.id}/pins/{message.id}")

	CrosspostMessage = NewAPIRoute(POST, "/channels/{channel.id}/messages/{message.id}/crosspost")

	GetReactions            = NewAPIRoute(GET, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}")
	AddReaction             = NewAPIRoute(PUT, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me")
	RemoveOwnReaction       = NewAPIRoute(DELETE, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me")
	RemoveUserReaction      = NewAPIRoute(DELETE, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/{user.id}")
	RemoveAllReactions      = NewAPIRoute(DELETE, "/channels/{channel.id}/messages/{message.id}/reactions")
	RemoveAllReactionsEmoji = NewAPIRoute(DELETE, "/channels/{channel.id}/messages/{message.id}/reactions/{emoji}")
)

// Emotes
var (
	GetEmotes   = NewAPIRoute(GET, "/guilds/{guild.id}/emojis")
	Getemote    = NewAPIRoute(GET, "/guilds/{guild.id}/emojis/{emoji.id}")
	CreateEmote = NewAPIRoute(POST, "/guilds/{guild.id}/emojis")
	UpdateEmote = NewAPIRoute(PATCH, "/guilds/{guild.id}/emojis/{emote.id}")
	DeleteEmote = NewAPIRoute(DELETE, "/guilds/{guild.id}/emojis/{emote.id}")
)

// Webhooks
var (
	GetWebhook             = NewAPIRoute(GET, "/webhooks/{webhook.id}")
	GetWebhookWithToken    = NewAPIRoute(GET, "/webhooks/{webhook.id}/{token}")
	UpdateWebhook          = NewAPIRoute(PATCH, "/webhooks/{webhook.id}")
	UpdateWebhokWithToken  = NewAPIRoute(PATCH, "/webhooks/{webhook.id}/{token}")
	DeleteWebhook          = NewAPIRoute(DELETE, "/webhooks/{webhook.id}")
	DeleteWebhookWithToken = NewAPIRoute(DELETE, "/webhooks/{webhook.id}/{token}")

	CreateWebhookMessage       = NewAPIRoute(POST, "/webhooks/{webhook.id}", "wait", "thread_id")
	CreateWebhookMessageSlack  = NewAPIRoute(POST, "/webhooks/{webhook.id}/slack", "wait")
	CreateWebhookMessageGithub = NewAPIRoute(POST, "/webhooks/{webhook.id}/github", "wait")
	UpdateWebhookMessage       = NewAPIRoute(POST, "/webhooks/{webhook.id}/{token}/messages/{message.id}")
	DeleteWebhookMessage       = NewAPIRoute(POST, "/webhooks/{webhook.id}/{token}/messages/{message.id}")
)

// Invites
var (
	GetInvite    = NewAPIRoute(GET, "/invites/{code}")
	CreateInvite = NewAPIRoute(POST, "/channels/{channel.id}/invites")
	DeleteInvite = NewAPIRoute(DELETE, "/invites/{code}")

	GetGuildInvite    = NewAPIRoute(GET, "/guilds/{guild.id}/invites")
	GetChannelInvites = NewAPIRoute(GET, "/channels/{channel.id}/invites")
)

// Interactions
var (
	GetGlobalCommands   = NewAPIRoute(GET, "/applications/{application.id}/commands")
	GetGlobalCommand    = NewAPIRoute(GET, "/applications/{application.id}/command/{command.id}")
	CreateGlobalCommand = NewAPIRoute(POST, "/applications/{application.id}/commands")
	SetGlobalCommands   = NewAPIRoute(PUT, "/applications/{application.id}/commands")
	UpdateGlobalCommand = NewAPIRoute(PATCH, "/applications/{application.id}/commands/{command.id}")
	DeleteGlobalCommand = NewAPIRoute(DELETE, "/applications/{application.id}/commands")

	GetGuildCommands   = NewAPIRoute(GET, "/applications/{application.id}/guilds/{guild.id}/commands")
	GetGuildCommand    = NewAPIRoute(GET, "/applications/{application.id}/guilds/{guild.id}/command/{command.id}")
	CreateGuildCommand = NewAPIRoute(POST, "/applications/{application.id}/guilds/{guild.id}/commands")
	SetGuildCommands   = NewAPIRoute(PUT, "/applications/{application.id}/guilds/{guild.id}/commands")
	UpdateGuildCommand = NewAPIRoute(PATCH, "/applications/{application.id}/guilds/{guild.id}/commands/{command.id}")
	DeleteGuildCommand = NewAPIRoute(DELETE, "/applications/{application.id}/guilds/{guild.id}/commands")

	GetGuildCommandPermissions  = NewAPIRoute(GET, "/applications/{application.id}/guilds/{guild.id}/commands/permissions")
	GetGuildCommandPermission   = NewAPIRoute(GET, "/applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions")
	SetGuildCommandsPermissions = NewAPIRoute(PUT, "/applications/{application.id}/guilds/{guild.id}/commands/permissions")
	SetGuildCommandPermissions  = NewAPIRoute(PUT, "/applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions")

	CreateInteractionResponse = NewAPIRoute(POST, "/interactions/{interaction.id}/{interaction.token}/callback")
	UpdateInteractionResponse = NewAPIRoute(PATCH, "/webhooks/{application.id}/{interaction.token}/messages/@original")
	DeleteInteractionResponse = NewAPIRoute(DELETE, "/webhooks/{application.id}/{interaction.token}/messages/@original")

	CreateFollowupMessage = NewAPIRoute(POST, "/webhooks/{application.id}/{interaction.token}")
	UpdateFollowupMessage = NewAPIRoute(PATCH, "/webhooks/{application.id}/{interaction.token}/messages/{message.id}")
	DeleteFollowupMessage = NewAPIRoute(DELETE, "/webhooks/{application.id}/{interaction.token}/messages/{message.id}")
)

// CDN
var (
	Emote                = NewCDNRoute("/emojis/{emote.id}.", []FileExtension{PNG, GIF})
	GuildIcon            = NewCDNRoute("/icons/{guild.id}/{icon.hash}.", []FileExtension{PNG, JPEG, WEBP, GIF})
	GuildSplash          = NewCDNRoute("/splashes/{guild.id}/guild.splash.", []FileExtension{PNG, JPEG, WEBP})
	GuildDiscoverySplash = NewCDNRoute("/discovery-splashes/{guild.id}/guild.discovery.splash.", []FileExtension{PNG, JPEG, WEBP})
	GuildBanner          = NewCDNRoute("/banners/{guild.id}/guild.banner.", []FileExtension{PNG, JPEG, WEBP})
	DefaultUserAvatar    = NewCDNRoute("/embed/avatars/{user.discriminator%5}.", []FileExtension{PNG})
	UserAvatar           = NewCDNRoute("/avatars/{user.id}/user.avatar.", []FileExtension{PNG, JPEG, WEBP, GIF})
	ApplicationIcon      = NewCDNRoute("/app-icons/{application.id}/icon.", []FileExtension{PNG, JPEG, WEBP})
	ApplicationAsset     = NewCDNRoute("/app-assets/{application.id}/{asset.id}.", []FileExtension{PNG, JPEG, WEBP})
	AchievementIcon      = NewCDNRoute("/app-assets/{application.id}/achievements/{achievement.id}/icons/{icon.hash}.", []FileExtension{PNG, JPEG, WEBP})
	TeamIcon             = NewCDNRoute("/team-icons/{team.id}/team.icon.", []FileExtension{PNG, JPEG, WEBP})
	Attachments          = NewCDNRoute("/attachments/{channel.id}/{attachment.id}/{file.name}", []FileExtension{BLANK})
)

// Other
var (
	InviteURL = NewRoute("https://discord.gg/{code}")
)
