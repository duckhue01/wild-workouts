module github.com/tribefintech/microservices/internal/notif

go 1.20

require (
	github.com/OneSignal/onesignal-go-api v1.0.4
	github.com/deepmap/oapi-codegen v1.12.4
	github.com/go-chi/render v1.0.2
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.9
	github.com/nats-io/nats.go v1.25.0
	github.com/sirupsen/logrus v1.9.2
	github.com/tabbed/pqtype v0.1.1
	github.com/tribefintech/microservices/internal/common v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.10.0
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-chi/chi/v5 v5.0.8
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nats-io/nats-server/v2 v2.9.17 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.15.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tribefintech/microservices/internal/common => ../common/
