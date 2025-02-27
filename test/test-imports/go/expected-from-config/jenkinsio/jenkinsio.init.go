package jenkinsio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"jenkinsio.Jenkins",
		reflect.TypeOf((*Jenkins)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Jenkins{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsProps",
		reflect.TypeOf((*JenkinsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpec",
		reflect.TypeOf((*JenkinsSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecBackup",
		reflect.TypeOf((*JenkinsSpecBackup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecBackupAction",
		reflect.TypeOf((*JenkinsSpecBackupAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecBackupActionExec",
		reflect.TypeOf((*JenkinsSpecBackupActionExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecConfigurationAsCode",
		reflect.TypeOf((*JenkinsSpecConfigurationAsCode)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecConfigurationAsCodeConfigurations",
		reflect.TypeOf((*JenkinsSpecConfigurationAsCodeConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecConfigurationAsCodeSecret",
		reflect.TypeOf((*JenkinsSpecConfigurationAsCodeSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecGroovyScripts",
		reflect.TypeOf((*JenkinsSpecGroovyScripts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecGroovyScriptsConfigurations",
		reflect.TypeOf((*JenkinsSpecGroovyScriptsConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecGroovyScriptsSecret",
		reflect.TypeOf((*JenkinsSpecGroovyScriptsSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecJenkinsApiSettings",
		reflect.TypeOf((*JenkinsSpecJenkinsApiSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMaster",
		reflect.TypeOf((*JenkinsSpecMaster)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterBasePlugins",
		reflect.TypeOf((*JenkinsSpecMasterBasePlugins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainers",
		reflect.TypeOf((*JenkinsSpecMasterContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnv",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvFrom",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvFromConfigMapRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvFromSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvValueFrom",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvValueFromFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*JenkinsSpecMasterContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecycle",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStart",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartExec",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStop",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopExec",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbe",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeExec",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeHttpGet",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersPorts",
		reflect.TypeOf((*JenkinsSpecMasterContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbe",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeExec",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeHttpGet",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"jenkinsio.JenkinsSpecMasterContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*JenkinsSpecMasterContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_JenkinsSpecMasterContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersResources",
		reflect.TypeOf((*JenkinsSpecMasterContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersSecurityContext",
		reflect.TypeOf((*JenkinsSpecMasterContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersSecurityContextCapabilities",
		reflect.TypeOf((*JenkinsSpecMasterContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*JenkinsSpecMasterContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*JenkinsSpecMasterContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterContainersVolumeMounts",
		reflect.TypeOf((*JenkinsSpecMasterContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterImagePullSecrets",
		reflect.TypeOf((*JenkinsSpecMasterImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterPlugins",
		reflect.TypeOf((*JenkinsSpecMasterPlugins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterSecurityContext",
		reflect.TypeOf((*JenkinsSpecMasterSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterSecurityContextSeLinuxOptions",
		reflect.TypeOf((*JenkinsSpecMasterSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterSecurityContextSysctls",
		reflect.TypeOf((*JenkinsSpecMasterSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterSecurityContextWindowsOptions",
		reflect.TypeOf((*JenkinsSpecMasterSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterTolerations",
		reflect.TypeOf((*JenkinsSpecMasterTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumes",
		reflect.TypeOf((*JenkinsSpecMasterVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesAwsElasticBlockStore",
		reflect.TypeOf((*JenkinsSpecMasterVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesAzureDisk",
		reflect.TypeOf((*JenkinsSpecMasterVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesAzureFile",
		reflect.TypeOf((*JenkinsSpecMasterVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCephfs",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCephfsSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCinder",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCinderSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesConfigMap",
		reflect.TypeOf((*JenkinsSpecMasterVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesConfigMapItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCsi",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesDownwardApi",
		reflect.TypeOf((*JenkinsSpecMasterVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesDownwardApiItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesEmptyDir",
		reflect.TypeOf((*JenkinsSpecMasterVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesFc",
		reflect.TypeOf((*JenkinsSpecMasterVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesFlexVolume",
		reflect.TypeOf((*JenkinsSpecMasterVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesFlocker",
		reflect.TypeOf((*JenkinsSpecMasterVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesGcePersistentDisk",
		reflect.TypeOf((*JenkinsSpecMasterVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesGitRepo",
		reflect.TypeOf((*JenkinsSpecMasterVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesGlusterfs",
		reflect.TypeOf((*JenkinsSpecMasterVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesHostPath",
		reflect.TypeOf((*JenkinsSpecMasterVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesIscsi",
		reflect.TypeOf((*JenkinsSpecMasterVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesIscsiSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesNfs",
		reflect.TypeOf((*JenkinsSpecMasterVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesPersistentVolumeClaim",
		reflect.TypeOf((*JenkinsSpecMasterVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesPhotonPersistentDisk",
		reflect.TypeOf((*JenkinsSpecMasterVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesPortworxVolume",
		reflect.TypeOf((*JenkinsSpecMasterVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjected",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSources",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesSecret",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*JenkinsSpecMasterVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesQuobyte",
		reflect.TypeOf((*JenkinsSpecMasterVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesRbd",
		reflect.TypeOf((*JenkinsSpecMasterVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesRbdSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesScaleIo",
		reflect.TypeOf((*JenkinsSpecMasterVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesScaleIoSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesSecret",
		reflect.TypeOf((*JenkinsSpecMasterVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesSecretItems",
		reflect.TypeOf((*JenkinsSpecMasterVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesStorageos",
		reflect.TypeOf((*JenkinsSpecMasterVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesStorageosSecretRef",
		reflect.TypeOf((*JenkinsSpecMasterVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecMasterVolumesVsphereVolume",
		reflect.TypeOf((*JenkinsSpecMasterVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotifications",
		reflect.TypeOf((*JenkinsSpecNotifications)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsMailgun",
		reflect.TypeOf((*JenkinsSpecNotificationsMailgun)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsMailgunApiKeySecretKeySelector",
		reflect.TypeOf((*JenkinsSpecNotificationsMailgunApiKeySecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsMailgunApiKeySecretKeySelectorSecret",
		reflect.TypeOf((*JenkinsSpecNotificationsMailgunApiKeySecretKeySelectorSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSlack",
		reflect.TypeOf((*JenkinsSpecNotificationsSlack)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSlackWebHookUrlSecretKeySelector",
		reflect.TypeOf((*JenkinsSpecNotificationsSlackWebHookUrlSecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSlackWebHookUrlSecretKeySelectorSecret",
		reflect.TypeOf((*JenkinsSpecNotificationsSlackWebHookUrlSecretKeySelectorSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSmtp",
		reflect.TypeOf((*JenkinsSpecNotificationsSmtp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSmtpPasswordSecretKeySelector",
		reflect.TypeOf((*JenkinsSpecNotificationsSmtpPasswordSecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSmtpPasswordSecretKeySelectorSecret",
		reflect.TypeOf((*JenkinsSpecNotificationsSmtpPasswordSecretKeySelectorSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSmtpUsernameSecretKeySelector",
		reflect.TypeOf((*JenkinsSpecNotificationsSmtpUsernameSecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsSmtpUsernameSecretKeySelectorSecret",
		reflect.TypeOf((*JenkinsSpecNotificationsSmtpUsernameSecretKeySelectorSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsTeams",
		reflect.TypeOf((*JenkinsSpecNotificationsTeams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsTeamsWebHookUrlSecretKeySelector",
		reflect.TypeOf((*JenkinsSpecNotificationsTeamsWebHookUrlSecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecNotificationsTeamsWebHookUrlSecretKeySelectorSecret",
		reflect.TypeOf((*JenkinsSpecNotificationsTeamsWebHookUrlSecretKeySelectorSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecRestore",
		reflect.TypeOf((*JenkinsSpecRestore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecRestoreAction",
		reflect.TypeOf((*JenkinsSpecRestoreAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecRestoreActionExec",
		reflect.TypeOf((*JenkinsSpecRestoreActionExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecRoles",
		reflect.TypeOf((*JenkinsSpecRoles)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecSeedJobs",
		reflect.TypeOf((*JenkinsSpecSeedJobs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecService",
		reflect.TypeOf((*JenkinsSpecService)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecServiceAccount",
		reflect.TypeOf((*JenkinsSpecServiceAccount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"jenkinsio.JenkinsSpecSlaveService",
		reflect.TypeOf((*JenkinsSpecSlaveService)(nil)).Elem(),
	)
}
