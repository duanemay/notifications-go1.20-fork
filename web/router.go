package web

import (
    "strings"

    "github.com/cloudfoundry-incubator/notifications/web/handlers"
    "github.com/gorilla/mux"
    "github.com/ryanmoran/stack"
)

type Router struct {
    stacks map[string]stack.Stack
}

func NewRouter(mother *Mother) Router {
    registrar := mother.Registrar()
    notify := handlers.NewNotify(mother.Courier(), mother.Finder(), registrar)
    preference := mother.Preference()
    preferenceUpdater := mother.PreferenceUpdater()
    logging := mother.Logging()
    errorWriter := mother.ErrorWriter()
    notificationsWriteAuthenticator := mother.Authenticator([]string{"notifications.write"})
    notificationPreferencesReadAuthenticator := mother.Authenticator([]string{"notification_preferences.read"})
    notificationPreferencesWriteAuthenticator := mother.Authenticator([]string{"notification_preferences.write"})

    return Router{
        stacks: map[string]stack.Stack{
            "GET /info":               stack.NewStack(handlers.NewGetInfo()).Use(logging),
            "POST /users/{guid}":      stack.NewStack(handlers.NewNotifyUser(notify, errorWriter)).Use(logging, notificationsWriteAuthenticator),
            "POST /spaces/{guid}":     stack.NewStack(handlers.NewNotifySpace(notify, errorWriter)).Use(logging, notificationsWriteAuthenticator),
            "PUT /registration":       stack.NewStack(handlers.NewRegistration(registrar, errorWriter)).Use(logging, notificationsWriteAuthenticator),
            "GET /user_preferences":   stack.NewStack(handlers.NewPreferenceFinder(preference, errorWriter)).Use(logging, notificationPreferencesReadAuthenticator),
            "PATCH /user_preferences": stack.NewStack(handlers.NewUpdatePreferences(preferenceUpdater, errorWriter)).Use(logging, notificationPreferencesWriteAuthenticator),
        },
    }
}

func (router Router) Routes() *mux.Router {
    r := mux.NewRouter()
    for methodPath, stack := range router.stacks {
        var name = methodPath
        parts := strings.SplitN(methodPath, " ", 2)
        r.Handle(parts[1], stack).Methods(parts[0]).Name(name)
    }
    return r
}
