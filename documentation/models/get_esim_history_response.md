# GetEsimHistoryResponse

**Properties**

| Name | Type                                                      | Required | Description |
| :--- | :-------------------------------------------------------- | :------- | :---------- |
| Esim | [GetEsimHistoryResponseEsim](#getesimhistoryresponseesim) | ✅        |             |

# GetEsimHistoryResponseEsim

**Properties**

| Name    | Type                                                                               | Required | Description    |
| :------ | :--------------------------------------------------------------------------------- | :------- | :------------- |
| Iccid   | string                                                                             | ✅        | ID of the eSIM |
| History | [[]*GetEsimHistoryResponseEsimHistoryItem](#getesimhistoryresponseesimhistoryitem) | ✅        |                |

# GetEsimHistoryResponseEsimHistoryItem

**Properties**

| Name       | Type      | Required | Description                                                                                                                         |
| :--------- | :-------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------- |
| Status     | string    | ✅        | The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR' |
| StatusDate | time.Time | ✅        | The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                         |
| Date       | float64   | ❌        | Epoch value representing the date when the eSIM status changed                                                                      |
