package manifest

import "encoding/xml"

type Manifest struct {
   XMLName                   xml.Name `xml:"manifest"`
   Text                      string   `xml:",chardata"`
   Android                   string   `xml:"android,attr"`
   VersionCode               string   `xml:"versionCode,attr"`
   VersionName               string   `xml:"versionName,attr"`
   CompileSdkVersion         string   `xml:"compileSdkVersion,attr"`
   CompileSdkVersionCodename string   `xml:"compileSdkVersionCodename,attr"`
   EmergencyInstaller        string   `xml:"emergencyInstaller,attr"`
   Package                   string   `xml:"package,attr"`
   PlatformBuildVersionCode  string   `xml:"platformBuildVersionCode,attr"`
   PlatformBuildVersionName  string   `xml:"platformBuildVersionName,attr"`
   UsesSdk                   struct {
      Text             string `xml:",chardata"`
      MinSdkVersion    string `xml:"minSdkVersion,attr"`
      TargetSdkVersion string `xml:"targetSdkVersion,attr"`
   } `xml:"uses-sdk"`
   UsesPermission []struct {
      Text                string `xml:",chardata"`
      Name                string `xml:"name,attr"`
      UsesPermissionFlags string `xml:"usesPermissionFlags,attr"`
      Required            string `xml:"required,attr"`
      MaxSdkVersion       string `xml:"maxSdkVersion,attr"`
   } `xml:"uses-permission"`
   UsesPermissionSdk23 struct {
      Text          string `xml:",chardata"`
      Name          string `xml:"name,attr"`
      MaxSdkVersion string `xml:"maxSdkVersion,attr"`
   } `xml:"uses-permission-sdk-23"`
   UsesFeature []struct {
      Text     string `xml:",chardata"`
      Name     string `xml:"name,attr"`
      Required string `xml:"required,attr"`
   } `xml:"uses-feature"`
   Permission []struct {
      Text            string `xml:",chardata"`
      Label           string `xml:"label,attr"`
      Name            string `xml:"name,attr"`
      ProtectionLevel string `xml:"protectionLevel,attr"`
      Description     string `xml:"description,attr"`
   } `xml:"permission"`
   Queries struct {
      Text    string `xml:",chardata"`
      Package struct {
         Text string `xml:",chardata"`
         Name string `xml:"name,attr"`
      } `xml:"package"`
      Provider struct {
         Text        string `xml:",chardata"`
         Name        string `xml:"name,attr"`
         Authorities string `xml:"authorities,attr"`
      } `xml:"provider"`
      Intent []struct {
         Text   string `xml:",chardata"`
         Action struct {
            Text string `xml:",chardata"`
            Name string `xml:"name,attr"`
         } `xml:"action"`
         Category struct {
            Text string `xml:",chardata"`
            Name string `xml:"name,attr"`
         } `xml:"category"`
         Data struct {
            Text   string `xml:",chardata"`
            Scheme string `xml:"scheme,attr"`
            Host   string `xml:"host,attr"`
         } `xml:"data"`
      } `xml:"intent"`
   } `xml:"queries"`
   Application struct {
      Text                         string `xml:",chardata"`
      Label                        string `xml:"label,attr"`
      Icon                         string `xml:"icon,attr"`
      Name                         string `xml:"name,attr"`
      BackupAgent                  string `xml:"backupAgent,attr"`
      KillAfterRestore             string `xml:"killAfterRestore,attr"`
      RestoreNeedsApplication      string `xml:"restoreNeedsApplication,attr"`
      RestoreAnyVersion            string `xml:"restoreAnyVersion,attr"`
      HardwareAccelerated          string `xml:"hardwareAccelerated,attr"`
      SupportsRtl                  string `xml:"supportsRtl,attr"`
      RequiredForAllUsers          string `xml:"requiredForAllUsers,attr"`
      RestrictedAccountType        string `xml:"restrictedAccountType,attr"`
      ExtractNativeLibs            string `xml:"extractNativeLibs,attr"`
      UsesCleartextTraffic         string `xml:"usesCleartextTraffic,attr"`
      RoundIcon                    string `xml:"roundIcon,attr"`
      AppComponentFactory          string `xml:"appComponentFactory,attr"`
      RequestLegacyExternalStorage string `xml:"requestLegacyExternalStorage,attr"`
      LocaleConfig                 string `xml:"localeConfig,attr"`
      EnableOnBackInvokedCallback  string `xml:"enableOnBackInvokedCallback,attr"`
      Service                      []struct {
         Text                  string `xml:",chardata"`
         Name                  string `xml:"name,attr"`
         Enabled               string `xml:"enabled,attr"`
         Exported              string `xml:"exported,attr"`
         Process               string `xml:"process,attr"`
         Permission            string `xml:"permission,attr"`
         ForegroundServiceType string `xml:"foregroundServiceType,attr"`
         VisibleToInstantApps  string `xml:"visibleToInstantApps,attr"`
         SingleUser            string `xml:"singleUser,attr"`
         IsolatedProcess       string `xml:"isolatedProcess,attr"`
         Description           string `xml:"description,attr"`
         DirectBootAware       string `xml:"directBootAware,attr"`
         IntentFilter          []struct {
            Text     string `xml:",chardata"`
            Priority string `xml:"priority,attr"`
            Action   []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Data struct {
               Text       string `xml:",chardata"`
               Scheme     string `xml:"scheme,attr"`
               Host       string `xml:"host,attr"`
               PathPrefix string `xml:"pathPrefix,attr"`
            } `xml:"data"`
            Category struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
         } `xml:"intent-filter"`
         MetaData []struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Resource string `xml:"resource,attr"`
            Value    string `xml:"value,attr"`
         } `xml:"meta-data"`
      } `xml:"service"`
      MetaData []struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Resource string `xml:"resource,attr"`
         Value    string `xml:"value,attr"`
      } `xml:"meta-data"`
      Activity []struct {
         Text                        string `xml:",chardata"`
         Theme                       string `xml:"theme,attr"`
         Name                        string `xml:"name,attr"`
         Exported                    string `xml:"exported,attr"`
         TaskAffinity                string `xml:"taskAffinity,attr"`
         ConfigChanges               string `xml:"configChanges,attr"`
         ExcludeFromRecents          string `xml:"excludeFromRecents,attr"`
         LaunchMode                  string `xml:"launchMode,attr"`
         ScreenOrientation           string `xml:"screenOrientation,attr"`
         NoHistory                   string `xml:"noHistory,attr"`
         Enabled                     string `xml:"enabled,attr"`
         Label                       string `xml:"label,attr"`
         WindowSoftInputMode         string `xml:"windowSoftInputMode,attr"`
         Icon                        string `xml:"icon,attr"`
         AllowTaskReparenting        string `xml:"allowTaskReparenting,attr"`
         Process                     string `xml:"process,attr"`
         Permission                  string `xml:"permission,attr"`
         VisibleToInstantApps        string `xml:"visibleToInstantApps,attr"`
         KnownActivityEmbeddingCerts string `xml:"knownActivityEmbeddingCerts,attr"`
         StateNotNeeded              string `xml:"stateNotNeeded,attr"`
         IntentFilter                []struct {
            Text       string `xml:",chardata"`
            Priority   string `xml:"priority,attr"`
            AutoVerify string `xml:"autoVerify,attr"`
            Action     []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Category []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
            Data []struct {
               Text       string `xml:",chardata"`
               MimeType   string `xml:"mimeType,attr"`
               Scheme     string `xml:"scheme,attr"`
               Host       string `xml:"host,attr"`
               PathPrefix string `xml:"pathPrefix,attr"`
               Path       string `xml:"path,attr"`
            } `xml:"data"`
         } `xml:"intent-filter"`
         MetaData []struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Resource string `xml:"resource,attr"`
            Value    string `xml:"value,attr"`
         } `xml:"meta-data"`
      } `xml:"activity"`
      Provider []struct {
         Text                 string `xml:",chardata"`
         Name                 string `xml:"name,attr"`
         Enabled              string `xml:"enabled,attr"`
         Exported             string `xml:"exported,attr"`
         Process              string `xml:"process,attr"`
         Authorities          string `xml:"authorities,attr"`
         GrantUriPermissions  string `xml:"grantUriPermissions,attr"`
         VisibleToInstantApps string `xml:"visibleToInstantApps,attr"`
         MetaData             []struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Value    string `xml:"value,attr"`
            Resource string `xml:"resource,attr"`
         } `xml:"meta-data"`
      } `xml:"provider"`
      Profileable struct {
         Text  string `xml:",chardata"`
         Shell string `xml:"shell,attr"`
      } `xml:"profileable"`
      UsesLibrary []struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Required string `xml:"required,attr"`
      } `xml:"uses-library"`
      Property []struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Value    string `xml:"value,attr"`
         Resource string `xml:"resource,attr"`
      } `xml:"property"`
      Receiver []struct {
         Text            string `xml:",chardata"`
         Name            string `xml:"name,attr"`
         Permission      string `xml:"permission,attr"`
         Exported        string `xml:"exported,attr"`
         Enabled         string `xml:"enabled,attr"`
         Process         string `xml:"process,attr"`
         DirectBootAware string `xml:"directBootAware,attr"`
         IntentFilter    []struct {
            Text   string `xml:",chardata"`
            Action []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Data struct {
               Text     string `xml:",chardata"`
               Scheme   string `xml:"scheme,attr"`
               MimeType string `xml:"mimeType,attr"`
            } `xml:"data"`
            Category []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
         } `xml:"intent-filter"`
         MetaData struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Resource string `xml:"resource,attr"`
         } `xml:"meta-data"`
      } `xml:"receiver"`
      ActivityAlias []struct {
         Text           string `xml:",chardata"`
         Name           string `xml:"name,attr"`
         Enabled        string `xml:"enabled,attr"`
         Exported       string `xml:"exported,attr"`
         TargetActivity string `xml:"targetActivity,attr"`
         Label          string `xml:"label,attr"`
         Icon           string `xml:"icon,attr"`
         RoundIcon      string `xml:"roundIcon,attr"`
         IntentFilter   struct {
            Text     string `xml:",chardata"`
            Category []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
            Action []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Data []struct {
               Text   string `xml:",chardata"`
               Scheme string `xml:"scheme,attr"`
               Host   string `xml:"host,attr"`
               Path   string `xml:"path,attr"`
            } `xml:"data"`
         } `xml:"intent-filter"`
         MetaData []struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Value    string `xml:"value,attr"`
            Resource string `xml:"resource,attr"`
         } `xml:"meta-data"`
      } `xml:"activity-alias"`
   } `xml:"application"`
} 
