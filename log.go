package gotel

// LogLevel Log level
type LogLevel int

const (
    // LogError Error level
    LogError LogLevel = iota
    
    // LogWarning Warning level
    LogWarning LogLevel = iota
    
    // LogDebug Debug level
    LogDebug LogLevel = iota
    
    // LogInfo Info level
    LogInfo LogLevel = iota
)