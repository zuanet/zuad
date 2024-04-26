package protowire

import (
	"github.com/zuanet/zuad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *ZuadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "ZuadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *ZuadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
