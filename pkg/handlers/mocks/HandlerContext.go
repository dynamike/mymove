// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/transcom/mymove/pkg/auth"

	dpsauth "github.com/transcom/mymove/pkg/dpsauth"

	handlers "github.com/transcom/mymove/pkg/handlers"

	http "net/http"

	iws "github.com/transcom/mymove/pkg/iws"

	mock "github.com/stretchr/testify/mock"

	notifications "github.com/transcom/mymove/pkg/notifications"

	pop "github.com/gobuffalo/pop/v5"

	route "github.com/transcom/mymove/pkg/route"

	scs "github.com/alexedwards/scs/v2"

	sequence "github.com/transcom/mymove/pkg/db/sequence"

	services "github.com/transcom/mymove/pkg/services"

	storage "github.com/transcom/mymove/pkg/storage"

	uuid "github.com/gofrs/uuid"
)

// HandlerContext is an autogenerated mock type for the HandlerContext type
type HandlerContext struct {
	mock.Mock
}

// AppNames provides a mock function with given fields:
func (_m *HandlerContext) AppNames() auth.ApplicationServername {
	ret := _m.Called()

	var r0 auth.ApplicationServername
	if rf, ok := ret.Get(0).(func() auth.ApplicationServername); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(auth.ApplicationServername)
	}

	return r0
}

// CookieSecret provides a mock function with given fields:
func (_m *HandlerContext) CookieSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *HandlerContext) DB() *pop.Connection {
	ret := _m.Called()

	var r0 *pop.Connection
	if rf, ok := ret.Get(0).(func() *pop.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pop.Connection)
		}
	}

	return r0
}

// DPSAuthParams provides a mock function with given fields:
func (_m *HandlerContext) DPSAuthParams() dpsauth.Params {
	ret := _m.Called()

	var r0 dpsauth.Params
	if rf, ok := ret.Get(0).(func() dpsauth.Params); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dpsauth.Params)
	}

	return r0
}

// FileStorer provides a mock function with given fields:
func (_m *HandlerContext) FileStorer() storage.FileStorer {
	ret := _m.Called()

	var r0 storage.FileStorer
	if rf, ok := ret.Get(0).(func() storage.FileStorer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.FileStorer)
		}
	}

	return r0
}

// GHCPlanner provides a mock function with given fields:
func (_m *HandlerContext) GHCPlanner() route.Planner {
	ret := _m.Called()

	var r0 route.Planner
	if rf, ok := ret.Get(0).(func() route.Planner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Planner)
		}
	}

	return r0
}

// GetFeatureFlag provides a mock function with given fields: name
func (_m *HandlerContext) GetFeatureFlag(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetTraceID provides a mock function with given fields:
func (_m *HandlerContext) GetTraceID() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// GexSender provides a mock function with given fields:
func (_m *HandlerContext) GexSender() services.GexSender {
	ret := _m.Called()

	var r0 services.GexSender
	if rf, ok := ret.Get(0).(func() services.GexSender); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.GexSender)
		}
	}

	return r0
}

// ICNSequencer provides a mock function with given fields:
func (_m *HandlerContext) ICNSequencer() sequence.Sequencer {
	ret := _m.Called()

	var r0 sequence.Sequencer
	if rf, ok := ret.Get(0).(func() sequence.Sequencer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sequence.Sequencer)
		}
	}

	return r0
}

// IWSPersonLookup provides a mock function with given fields:
func (_m *HandlerContext) IWSPersonLookup() iws.PersonLookup {
	ret := _m.Called()

	var r0 iws.PersonLookup
	if rf, ok := ret.Get(0).(func() iws.PersonLookup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iws.PersonLookup)
		}
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *HandlerContext) Logger() handlers.Logger {
	ret := _m.Called()

	var r0 handlers.Logger
	if rf, ok := ret.Get(0).(func() handlers.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handlers.Logger)
		}
	}

	return r0
}

// LoggerFromContext provides a mock function with given fields: ctx
func (_m *HandlerContext) LoggerFromContext(ctx context.Context) handlers.Logger {
	ret := _m.Called(ctx)

	var r0 handlers.Logger
	if rf, ok := ret.Get(0).(func(context.Context) handlers.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handlers.Logger)
		}
	}

	return r0
}

// LoggerFromRequest provides a mock function with given fields: r
func (_m *HandlerContext) LoggerFromRequest(r *http.Request) handlers.Logger {
	ret := _m.Called(r)

	var r0 handlers.Logger
	if rf, ok := ret.Get(0).(func(*http.Request) handlers.Logger); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handlers.Logger)
		}
	}

	return r0
}

// NotificationSender provides a mock function with given fields:
func (_m *HandlerContext) NotificationSender() notifications.NotificationSender {
	ret := _m.Called()

	var r0 notifications.NotificationSender
	if rf, ok := ret.Get(0).(func() notifications.NotificationSender); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(notifications.NotificationSender)
		}
	}

	return r0
}

// Planner provides a mock function with given fields:
func (_m *HandlerContext) Planner() route.Planner {
	ret := _m.Called()

	var r0 route.Planner
	if rf, ok := ret.Get(0).(func() route.Planner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Planner)
		}
	}

	return r0
}

// SendProductionInvoice provides a mock function with given fields:
func (_m *HandlerContext) SendProductionInvoice() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SessionAndLoggerFromContext provides a mock function with given fields: ctx
func (_m *HandlerContext) SessionAndLoggerFromContext(ctx context.Context) (*auth.Session, handlers.Logger) {
	ret := _m.Called(ctx)

	var r0 *auth.Session
	if rf, ok := ret.Get(0).(func(context.Context) *auth.Session); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Session)
		}
	}

	var r1 handlers.Logger
	if rf, ok := ret.Get(1).(func(context.Context) handlers.Logger); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(handlers.Logger)
		}
	}

	return r0, r1
}

// SessionAndLoggerFromRequest provides a mock function with given fields: r
func (_m *HandlerContext) SessionAndLoggerFromRequest(r *http.Request) (*auth.Session, handlers.Logger) {
	ret := _m.Called(r)

	var r0 *auth.Session
	if rf, ok := ret.Get(0).(func(*http.Request) *auth.Session); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Session)
		}
	}

	var r1 handlers.Logger
	if rf, ok := ret.Get(1).(func(*http.Request) handlers.Logger); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(handlers.Logger)
		}
	}

	return r0, r1
}

// SessionFromContext provides a mock function with given fields: ctx
func (_m *HandlerContext) SessionFromContext(ctx context.Context) *auth.Session {
	ret := _m.Called(ctx)

	var r0 *auth.Session
	if rf, ok := ret.Get(0).(func(context.Context) *auth.Session); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Session)
		}
	}

	return r0
}

// SessionFromRequest provides a mock function with given fields: r
func (_m *HandlerContext) SessionFromRequest(r *http.Request) *auth.Session {
	ret := _m.Called(r)

	var r0 *auth.Session
	if rf, ok := ret.Get(0).(func(*http.Request) *auth.Session); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Session)
		}
	}

	return r0
}

// SessionManager provides a mock function with given fields: session
func (_m *HandlerContext) SessionManager(session *auth.Session) *scs.SessionManager {
	ret := _m.Called(session)

	var r0 *scs.SessionManager
	if rf, ok := ret.Get(0).(func(*auth.Session) *scs.SessionManager); ok {
		r0 = rf(session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scs.SessionManager)
		}
	}

	return r0
}

// SetAppNames provides a mock function with given fields: appNames
func (_m *HandlerContext) SetAppNames(appNames auth.ApplicationServername) {
	_m.Called(appNames)
}

// SetCookieSecret provides a mock function with given fields: secret
func (_m *HandlerContext) SetCookieSecret(secret string) {
	_m.Called(secret)
}

// SetDPSAuthParams provides a mock function with given fields: params
func (_m *HandlerContext) SetDPSAuthParams(params dpsauth.Params) {
	_m.Called(params)
}

// SetFeatureFlag provides a mock function with given fields: flags
func (_m *HandlerContext) SetFeatureFlag(flags handlers.FeatureFlag) {
	_m.Called(flags)
}

// SetFileStorer provides a mock function with given fields: storer
func (_m *HandlerContext) SetFileStorer(storer storage.FileStorer) {
	_m.Called(storer)
}

// SetGHCPlanner provides a mock function with given fields: planner
func (_m *HandlerContext) SetGHCPlanner(planner route.Planner) {
	_m.Called(planner)
}

// SetGexSender provides a mock function with given fields: gexSender
func (_m *HandlerContext) SetGexSender(gexSender services.GexSender) {
	_m.Called(gexSender)
}

// SetICNSequencer provides a mock function with given fields: sequencer
func (_m *HandlerContext) SetICNSequencer(sequencer sequence.Sequencer) {
	_m.Called(sequencer)
}

// SetIWSPersonLookup provides a mock function with given fields: rbs
func (_m *HandlerContext) SetIWSPersonLookup(rbs iws.PersonLookup) {
	_m.Called(rbs)
}

// SetNotificationSender provides a mock function with given fields: sender
func (_m *HandlerContext) SetNotificationSender(sender notifications.NotificationSender) {
	_m.Called(sender)
}

// SetPlanner provides a mock function with given fields: planner
func (_m *HandlerContext) SetPlanner(planner route.Planner) {
	_m.Called(planner)
}

// SetSendProductionInvoice provides a mock function with given fields: sendProductionInvoice
func (_m *HandlerContext) SetSendProductionInvoice(sendProductionInvoice bool) {
	_m.Called(sendProductionInvoice)
}

// SetSessionManagers provides a mock function with given fields: sessionManagers
func (_m *HandlerContext) SetSessionManagers(sessionManagers [3]*scs.SessionManager) {
	_m.Called(sessionManagers)
}

// SetTraceID provides a mock function with given fields: traceID
func (_m *HandlerContext) SetTraceID(traceID uuid.UUID) {
	_m.Called(traceID)
}

// SetUseSecureCookie provides a mock function with given fields: useSecureCookie
func (_m *HandlerContext) SetUseSecureCookie(useSecureCookie bool) {
	_m.Called(useSecureCookie)
}

// UseSecureCookie provides a mock function with given fields:
func (_m *HandlerContext) UseSecureCookie() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
