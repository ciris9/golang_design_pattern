package main

// Button 请求者
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Command 命令接口
type Command interface {
	execute()
}

// OnCommand 具体命令1
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// OffCommand 具体命令2
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Device 接受者接口
type Device interface {
	on()
	off()
}

// Tv 具体接受者1
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
}

func (t *Tv) off() {
	t.isRunning = false
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
