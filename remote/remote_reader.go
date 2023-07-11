package remote

import (
	context "context"
	"fmt"
	"main/proto"

	"github.com/google/uuid"
)

type RemoteReader struct {
	proto.UnimplementedRemoteServer
	remote *Remote
}

func NewRemoteReader(remote *Remote) *RemoteReader {
	return &RemoteReader{
		remote: remote,
	}
}

func (r *RemoteReader) GetRemotingActor(ctx context.Context, request *proto.RemotingActorRequest) (*proto.RemotingActorResponse, error) {
	pid := r.remote.system.ActorPidByName(request.Name)

	if pid == uuid.Nil {
		return &proto.RemotingActorResponse{
			Pid: uuid.Nil.String(),
		}, fmt.Errorf("Actor of name %s is not found \n", request.Name)
	}
	return &proto.RemotingActorResponse{
		Pid: pid.String(),
	}, nil
}

func (r *RemoteReader) SendMessage(ctx context.Context, request *proto.ProtoEnvelope) (*proto.MessageResponse, error) {
	r.remote.system.Root.Send(uuid.MustParse(request.Target), proto.Decode(request))
	return &proto.MessageResponse{Message: "Remote system received the message \n"}, nil
}
