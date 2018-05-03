package services

type HeartMessage struct{
    Serverid  int
}


type HeartBeat int

func (hb *HeartBeat)HeartBeat(args *HeartMessage, reply *bool) error{
    *reply=true
    return nil
}
