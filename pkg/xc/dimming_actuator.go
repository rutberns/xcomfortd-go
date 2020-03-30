package xc

import "context"

/* New dimming actuator output channels:

   0 = status
   1 = binary input A
   2 = binary input B
   3 = energy
   4 = power
   5 = load error */

const (
	CDAU_0104   int = 0
	CDAU_0104_I     = 1 // 2 x binary input
	CDAU_0104_E     = 2 // Energy function
	CDAE_0104       = 4
	CDAE_0104_E     = 6  // Energy function
	CDAE_0105_I     = 9  // 2 x binary input
	CDAE_0105_E     = 10 // Energy function
)

var dimmerNames = map[int]string{
	CDAU_0104:   "CDAU 01/04",
	CDAU_0104_I: "CDAU 01/04-I",
	CDAU_0104_E: "CDAU 01/04-E",
	CDAE_0104:   "CDAE 01/04",
	CDAE_0104_E: "CDAE 01/04-E",
	CDAE_0105_I: "CDAE 01/05-I",
	CDAE_0105_E: "CDAE 01/05-E",
}

func (d *Datapoint) Dim(ctx context.Context, value int) ([]byte, error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.device.iface.sendTxCommand(ctx, []byte{d.number, TX_EVENT_DIM, TX_EVENTDATA_PERCENT, byte(value)})
}