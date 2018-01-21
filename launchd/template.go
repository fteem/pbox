package launchd

func PlistTemplate() string {
	return `
<?xml version='1.0' encoding='UTF-8'?>
<!DOCTYPE plist PUBLIC \"-//Apple Computer//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\" >
<plist version='1.0'>
  <dict>
    <key>Label</key><string>{{ .DisplayName }}</string>
    <key>Program</key><string>{{ .Program }}</string>
    {{ if .ProgramArguments }}
    <key>ProgramArguments</key>
    <array>
      {{ range $argument := .ProgramArguments }}
	<string>{{ $argument }}
      {{ end }}
    </array>
    {{ end }}
    <key>WorkingDirectory</key><string>{{ .WorkingDirectory }}</string>
    <key>StandardOutPath</key><string>{{ .LogLocation }}/{{ .DisplayName }}.log</string>
    <key>StandardErrorPath</key><string>{{ .LogLocation }}/{{ .DisplayName }}.err</string>
    {{ if .KeepAlive }}<key>KeepAlive</key><{{ .KeepAlive }}/>{{ end }}
    {{ if .Disabled }}<key>Disabled</key><{{ .Disabled }}/>{{ end }}
    {{ if .RunAtLoad }}<key>RunAtLoad</key><{{ .RunAtLoad }}/>{{ end }}
    {{ if .ThrottleInterval }}<key>ThrottleInterval</key><integer>{{ .ThrottleInterval }}</integer>{{ end }}
    {{ if .UserName }}<key>UserName</key><string>{{ .UserName }}</string>{{ end }}
    {{ if .StartInterval }}<key>StartInterval</key><integer>{{ .StartInterval }}</integer>{{ end }}
    {{ if .StartOnMount }}<key>StartOnMount</key><{{ .StartOnMount }}/>{{ end }}
    {{ if .GroupName }}<key>GroupName</key><string>{{ .GroupName }}</string>{{ end }}
    {{ if .InitGroups }}<key>InitGroups</key><{{ .InitGroups }}/>{{ end }}
    {{ if .ThrottleInterval }}<key>ThrottleInterval</key><integer>{{ .ThrottleInterval }}</integer>{{ end }}
    {{ if .WatchPaths }}
      <key>WatchPaths</key>
      <array>
	{{ range $path := .WatchPaths }}
          <string>{{ $path }}</string>
        {{ end }}
      </array>
    {{ end }}
  </dict>
</plist>
`
}
