// SiGG-GoLang-On-the-Fly //
package core

type LocallySequenced interface {
	LocalSequence() int64
}
