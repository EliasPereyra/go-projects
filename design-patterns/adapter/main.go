package main

func main() {
	client := &Client{}
	mac := &Mac{}

	// el cliente se conecta a la mac sin problemas
	client.InsertIntoLightningConnectorIntoComputer(mac)

	// el cliente se quiere conectar con el mismo dispositivo
	// a una maquina windows, pero no tiene la misma interfaz.
	// por lo que se necesitara de un adaptador para poder utilizarlo.
	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowMachine: windows,
	}

	// Ahora sí el cliente se puede conectar a una maquina windows
	// y tener la misma funcionalidad que en mac, a través del ".ter/client
	client.InsertIntoLightningConnectorIntoComputer(windowsAdapter)
}
