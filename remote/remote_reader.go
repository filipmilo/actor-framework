package remote

import (
	context "context"
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
	return &proto.RemotingActorResponse{
		Pid: r.remote.system.ActorPidByName(request.Name).String(),
	}, nil
}

func (r *RemoteReader) SendMessage(ctx context.Context, request *proto.ProtoEnvelope) (*proto.MessageResponse, error) {
	r.remote.system.Root.Send(uuid.MustParse(request.Target), proto.Decode(request))
	return &proto.MessageResponse{Message: "Remote system received the message"}, nil
}
