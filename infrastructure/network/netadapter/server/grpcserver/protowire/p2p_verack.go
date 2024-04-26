package protowire

import (
	"github.com/zuanet/zuad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *ZuadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "ZuadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *ZuadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
