package client

import (
	"io"
	pb "github.com/Azure/azure-functions-go-worker/internal/rpc"
)

type Client {
	stream.FunctionRpc_EventStreamClient
}

func New(stream pb.FunctionRpc_EventStreamClient) *Client{
	return &Client{
		stream: stream,
	}
}

func (c *Client) Listen(requestId, workerId string) error {
	c.stream.Send(&pb.StreamingMessage{
		RequestId: requestId,
		Content: &pb.StreaminMessage_StartStream{
			WorkerId: workerId,
		},
	})


	for {
		message, err := c.stream.Recv()

		if err == io.EOF {
			return nil
		}
		
		if err != nil {
			return err
		}


		if err := c.handle(requestId, message); err != nil {
			return err
		}
	}


	return nil
}


func (c *Client) handle(requestId string, message interface{}) error {

	switch msg := message.(type) {
	case *pb.StreamingMessage_WorkerInitRequest:
		return c.stream.Send(&pb.StreamingMessage{
			RequestId: requestId,
			Content: &pb.StreamingMessage_WorkerInitResponse{
				Result: &pb.StatusResult{
					Status: pb.StatusResult_Success,
				}
			},
		})	
	case *pb.StreamingMessage_FunctionLoadRequest:

		// TODO: This is a stub
		return c.stream.Send(&pb.StreamingMessage{
			RequestId: requestId,
			Content: &pb.StreamingMessage_FunctionLoadResponse{
				FunctionId: msg.FunctionId,
				Result: &pb.StatusResult{
					Status: pb.StatusResult_Success,
				},
			},
		})	
	case *pb.StreamingMessage_InvocationRequest:

		
	}
}