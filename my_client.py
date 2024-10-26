import grpc
import auth_pb2
import auth_pb2_grpc

def run():
    try:
        channel = grpc.insecure_channel('localhost:50052')
        stub = auth_pb2_grpc.AuthServiceStub(channel)
        
        response = stub.Ping(auth_pb2.PingRequest(client_id="test_client"))
        print("Ping response:", response.message)

    except grpc.RpcError as e:
        print(f"RPC failed with status {e.code()}: {e.details()}")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == '__main__':
    run()

