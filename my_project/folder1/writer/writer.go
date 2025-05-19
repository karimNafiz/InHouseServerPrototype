package writer


/*
<summary>
	This interface is meant to define behaviour for any writer that 
	will write to an IO, it could either be File IO or socket
	It will have a Write method that will take in the no of bytes to write
	return the no of bytes actually written and the error
</summary>
*/

type Writer interface {
	Write(bytes []byte) (int , error)
}