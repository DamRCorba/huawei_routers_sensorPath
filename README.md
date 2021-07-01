# Huawei Telemetry Sensors

This project is an update and rename of Huawei Sensor Path. Its a new repository in order to be compatible with Influx requirements. This version will support Huawei telemetry sensors for Huawei Routers & Switches running V8R10 to V8R12
## Telemetry Sensors V8R10

```
<Huawei>dis telemetry sensor-path | i Sample
Info: It will take a long time if the content you search is too much or the string you input is too long, you can press CTRL_C to break.
Sample(S) : Serial sample.
Sample(P) : Parallel sample.
Parallel sampled data may not be sent through UDP in out-of-band mode.
------------------------------------------------------------------------------------------------------------------
Type        MinPeriod(ms)  MaxEachPeriod  Path      
------------------------------------------------------------------------------------------------------------------
Sample(S)   10000          --             huawei-devm:devm/cpuInfos/cpuInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/fans/fan - Ok
Sample(S)   10000          --             huawei-devm:devm/memoryInfos/memoryInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/ports/port - Solo devuelve posicion y entIndex
Sample(S)   300000         --             huawei-devm:devm/ports/port/opticalInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/powerSupplys/powerSupply/powerEnvironments/powerEnvironment - Ok.
Sample(S)   300000         --             huawei-devm:devm/temperatureInfos/temperatureInfo - Ok
Sample(P)   300000         --             huawei-emdi:emdi/emdiTelemReps/emdiTelemRep -
Sample(P)   10000          --             huawei-emdi:emdi/emdiTelemRtps/emdiTelemRtp -
Sample(S)   10000          --             huawei-ifm:ifm/interfaces/interface - Solo devuelve AdminStatus si la interface esta UP. Sino solo index
Sample(P)   100            20             huawei-ifm:ifm/interfaces/interface/ifClearedStat - No Interesa
Sample(S)   10000          --             huawei-ifm:ifm/interfaces/interface/ifDynamicInfo - Devuelve el estado operativo de la interfaz
Sample(P)   100            20             huawei-ifm:ifm/interfaces/interface/ifStatistics - Probado ok
Sample(P)   1000           20             huawei-ifm:ifm/interfaces/interface/ifStatistics/ethPortErrSts - OK
Sample(S)   10000          --             huawei-mpls:mpls/mplsTe/rsvpTeTunnels/rsvpTeTunnel/tunnelPaths/tunnelPath - No hay RSVP
Sample(P)   1000           200            huawei-qos:qos/qosBuffers/qosBuffer - No da Informacion de metricas.
Sample(P)   100            20             huawei-qos:qos/qosIfQoss/qosIfQos/qosPolicyApplys/qosPolicyApply/qosPolicyStats/qosPolicyStat/qosRuleStats/qosRuleStat Da informacion de buffer por politica por puerto.
Sample(P)   1000           10             huawei-qos:qos/qosPortQueueStatInfos/qosPortQueueStatInfo - Probado Ok
Sample(S)   300000         --             huawei-trafficmng:trafficmng/tmSlotSFUs/tmSlotSFU/sfuStatisticss/sfuStatistics Probado Ok.
------------------------------------------------------------------------------------------------------------------

```

### V8R10 Sensor ISSUE

The NE40-X3/8/16A in VRP V8R10 sends the information of the sensor "huawei-qos:qos/qosPortQueueStatInfos/qosPortQueueStatInfo" misordered. The plugin not correct this input for future compatibility. Thats an ISSUE from Huawei so this is the information needed to use this plugin. Check if your device sends it correctly.


| Sensor Queue Indication | Real Queue     |
| :---------------------- | :------------- |
| Null                    | BE             |
| AF2                     | AF1            |
| AF3                     | AF2            |
| AF4                     | AF3            |
| BE                      | AF4            |
| CS6                     | EF             |
| CS7                     | CS6            |
| EF                      | CS7            |

## Telemetry Sensors V8R11/V8R12

```bash
------------------------------------------------------------------------------------------------------------------
Type          MinPeriod(ms)  MaxEachPeriod  Path      
------------------------------------------------------------------------------------------------------------------
Sample        1000           --             huawei-debug:debug/cpu-infos/cpu-info
Sample        1000           --             huawei-debug:debug/memory-infos/memory-info
Sample        300000         --             huawei-devm:devm/ports/port/huawei-pic:optical-module
Sample        300000         --             huawei-emdi:emdi/emdi-telem-reps/emdi-telem-rep
Sample        10000          --             huawei-emdi:emdi/emdi-telem-rtps/emdi-telem-rtp
Sample        300000         --             huawei-emdi:emdi/out-telem-reps/out-telem-rep
OnChange      --             --             huawei-esqm:esqm/gtp-datas/gtp-data
OnChange      --             --             huawei-esqm:esqm/tcp-datas/tcp-data
Sample        1000           20             huawei-flow-recognition:flow-recognition/streaminfos/streaminfo
OnChange      --             --             huawei-ifit:ifit/huawei-ifit-statistics:flow-hop-statistics/flow-hop-statistic
OnChange      --             --             huawei-ifit:ifit/huawei-ifit-statistics:flow-locator-statistics/flow-locator-statistic
OnChange      --             --             huawei-ifit:ifit/huawei-ifit-statistics:flow-peer-ip-statistics/flow-peer-ip-statistic
OnChange      --             --             huawei-ifit:ifit/huawei-ifit-statistics:flow-statistics/flow-statistic
Sample        1000           --             huawei-ifm:ifm/interfaces/interface
Sample        100            20             huawei-ifm:ifm/interfaces/interface/common-statistics
Sample        1000           --             huawei-ifm:ifm/interfaces/interface/dynamic
Sample        10             100            huawei-ifm:ifm/interfaces/interface/mib-statistics
Sample        1000           20             huawei-ifm:ifm/interfaces/interface/mib-statistics/huawei-pic:eth-port-err-sts
OnChange      --             --             huawei-kpi:kpi/kpi-datas/kpi-data
Sample        300000         --             huawei-pic:pic/port-statistics/statistic
Alarm         --             --             huawei-qos-notification:xqos-port-queue-alarm
Alarm         --             --             huawei-qos-notification:xqos-port-queue-alarm-resume
Sample        1000           1              huawei-qos:qos/global-query/all-queue-statisticss/all-queue-statistics
Sample        1000           10             huawei-qos:qos/global-query/channel-queue-statisticss/channel-queue-statistics
Sample        1000           10             huawei-qos:qos/global-query/default-queue-statisticss/default-queue-statistics
Sample        100            20             huawei-qos:qos/global-query/interface-traffic-policy-statisticss/interface-traffic-policy-statistics/classifier-based-staticss/classifier-based-statics
Sample        100            20             huawei-qos:qos/global-query/interface-traffic-policy-statisticss/interface-traffic-policy-statistics/rule-based-staticss/rule-based-statics
Sample        300000         --             huawei-trafficmng:trafficmng/tm-sfu-informations/tm-sfu-information/sfu-statistics/sfu-statistic
Sample        5000           512            huawei-twamp-controller:twamp-controller/client/sessions/session/huawei-twamp-statistics:statistics
Sample        1000           120            other paths
------------------------------------------------------------------------------------------------------------------
note:
1. Recommended configuration: Sampling period = (Total number of sampling instances) / (Number of sampling instances in each minimum sampling period) * (Minimum sampling period).
2. MaxEachPeriod indicates the historical maximum value, not the specification. 

```