package allRoutes

const (
	ROUT_TEST = "test"

	//----====WEBSOCKET====----\\

	ROUT_WEBSOCKET_CHAT_AND_ONLINE = "websocketChat"
	ROUT_WEBSOCKET_SOCIAL          = "websocketSocial"
	ROUT_WEBSOCKET_NOTIFICATION    = "websocketNotification"

	//----====MAIN====----\\

	ROUT_INDEX              = "index"
	ROUT_NEWS_VIEW          = "newsView"
	ROUT_ABOUT_GAME         = "aboutGame"
	ROUT_HISTORY            = "history"
	ROUT_RULES              = "rules"
	ROUT_AGREEMENT          = "agreement"
	ROUT_CONTACTS           = "contacts"
	ROUT_GET_SETTINGS_AUDIO = "getSettingsAudio"

	//----====USER====----\\

	ROUT_LOGIN               = "login"
	ROUT_SOCIAL              = "social"
	ROUT_ACCOUNT_RECOVERY    = "accountRecovery"
	ROUT_REGISTER            = "register"
	ROUT_LOGOUT              = "logout"
	ROUT_ACCOUNT             = "account"
	ROUT_COMPLAINT           = "complaint"
	ROUT_VIP_STATUS          = "vipStatus"
	ROUT_SQUAD               = "squad"
	ROUT_SQUAD_INVITE        = "squadInvite"
	ROUT_USERS_LIST_IGNORE   = "usersListIgnore"
	ROUT_USERS_LIST_OUTGOING = "usersListOutgoing"
	ROUT_USERS_LIST_INCOMING = "usersListIncoming"
	ROUT_USERS_LIST_FRIENDS  = "usersListFriends"
	ROUT_USERS_SEARCH        = "usersSearch"
	ROUT_MAIN_MENU           = "mainMenu"
	ROUT_SETTINGS_CONTROL    = "settingsControl"
	ROUT_SETTINGS_AUDIO      = "settingsAudio"
	ROUT_STATISTICS_GENERAL  = "statisticsGeneral"

	//----====LOBBY====----\\

	ROUT_LOBBY_CLASSES   = "lobbyClasses"
	ROUT_BUY_SKILL       = "buySkill"
	ROUT_BUY_MODEL       = "buyModel"
	ROUT_LOBBY_LOCATIONS = "lobbyLocations"
	ROUT_SAVE_CLASS      = "saveClass"
	ROUT_CHOICE_LOCATION = "choiceLocation"

	//----====FEEDBACK====----\\

	ROUT_FAQ                      = "faq"
	ROUT_TECH_SUPPORT_THEMES_LIST = "techSupportThemesList"
	ROUT_TECH_SUPPORT_THEME_VIEW  = "techSupportThemeView"
	ROUT_TECH_SUPPORT_THEME_NEW   = "techSupportThemeNew"

	//----====ADMIN====----\\
	ROUT_ADMIN_INDEX                     = "adminIndex"
	ROUT_ADMIN_ACCESS_REGISTER_LIST      = "adminAccessRegisterList"
	ROUT_ADMIN_ACCESS_REGISTER_VIEW      = "adminAccessRegisterView"
	ROUT_ADMIN_APPLICATION_REGISTER_LIST = "adminApplicationRegisterList"
	ROUT_ADMIN_APPLICATION_REGISTER_VIEW = "adminApplicationRegisterView"
	ROUT_ADMIN_USERS_LIST                = "adminUsersList"
	ROUT_ADMIN_USER_VIEW                 = "adminUserView"
	ROUT_ADMIN_WARNINGS_LIST             = "adminWarningsList"
	ROUT_ADMIN_WARNED_VIEW               = "adminWarnedView"
	ROUT_ADMIN_CLASSES_LIST              = "adminClassesList"
	ROUT_ADMIN_CLASS_VIEW                = "adminClassView"
	ROUT_ADMIN_STATISTICS_LIST           = "adminStatisticsList"
	ROUT_ADMIN_STATISTICS_VIEW           = "adminStatisticsView"
	ROUT_ADMIN_NOTIFICATIONS_LIST        = "adminNotificationsList"
	ROUT_ADMIN_NOTIFICATION_VIEW         = "adminNotificationView"
	ROUT_ADMIN_FAQ_LIST                  = "adminFaqList"
	ROUT_ADMIN_FAQ_VIEW                  = "adminFaqView"
	ROUT_ADMIN_TECH_SUPPORT_THEMES_LIST  = "adminTechSupportThemesList"
	ROUT_ADMIN_TECH_SUPPORT_THEME_VIEW   = "adminTechSupportThemeView"
	ROUT_ADMIN_CARTOGRAPHER              = "adminCartographer"
	ROUT_ADMIN_CARTOGRAPHER_SAVE         = "adminCartographerSave"
	ROUT_ADMIN_CARTOGRAPHER_GET_FILES    = "adminCartographerGetFiles"
	ROUT_ADMIN_CARTOGRAPHER_LOAD         = "adminCartographerLoad"
	ROUT_ADMIN_NEWS_LIST                 = "adminNewsList"
	ROUT_ADMIN_NEWS_VIEW                 = "adminNewsView"
	ROUT_ADMIN_COMPLAINTS_LIST           = "adminComplaintsList"
	ROUT_ADMIN_COMPLAINT_VIEW            = "adminComplaintView"
	ROUT_ADMIN_LOCATIONS_LIST            = "adminLocationsList"
	ROUT_ADMIN_LOCATION_VIEW             = "adminLocationView"
	ROUT_ADMIN_MESSAGES_CHAT_LIST        = "adminMessagesChatList"
	ROUT_ADMIN_MESSAGE_CHAT_VIEW         = "adminMessageChatView"
)

var Url map[string]string
