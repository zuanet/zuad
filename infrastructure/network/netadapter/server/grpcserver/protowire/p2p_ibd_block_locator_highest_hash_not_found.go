package protowire

import (
	"github.com/zuanet/zuad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *ZuadMessage_IbdBlockLocatorHighestHashNotFound) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "ZuadMessage_IbdBlockLocatorHighestHashNotFound is nil")
	}
	return &appmessage.MsgIBDBlockLocatorHighestHashNotFound{}, nil
}

func (x *ZuadMessage_IbdBlockLocatorHighestHashNotFound) fromAppMessage(message *appmessage.MsgIBDBlockLocatorHighestHashNotFound) error {
	x.IbdBlockLocatorHighestHashNotFound = &IbdBlockLocatorHighestHashNotFoundMessage{}
	return nil
}
