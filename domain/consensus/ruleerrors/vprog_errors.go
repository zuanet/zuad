package ruleerrors

var (
    ErrVProgCodeTooLarge        = errors.New("vprog code too large")
    ErrVProgDataTooLarge        = errors.New("vprog data too large") 
    ErrUnsupportedVProgVersion  = errors.New("unsupported vprog version")
    ErrVProgGasLimitExceeded    = errors.New("vprog gas limit exceeded")
    ErrVProgExecutionFailed     = errors.New("vprog execution failed")
    ErrVProgInvalidOpcode       = errors.New("invalid vprog opcode")
    ErrVProgOutOfGas            = errors.New("vprog out of gas")
)
