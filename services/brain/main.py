import grpc
from concurrent import futures
import time

import study_pb2
import study_pb2_grpc

class StudyEngineServicer(study_pb2_grpc.StudyEngineServicer):
    def SearchNotes(self, request, context):
        print(f"Python received search query: {request.query}")
        
        # Simulated AI response for now
        match = study_pb2.SearchResponse.Match(
            text=f"Found a relevant snippet for '{request.query}' in your notes.",
            source="notes/example.md",
            score=0.98
        )
        return study_pb2.SearchResponse(matches=[match])

    def IndexDocument(self, request, context):
        print(f"Python indexing file: {request.file_path}")
        return study_pb2.IndexResponse(success=True)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    study_pb2_grpc.add_StudyEngineServicer_to_server(StudyEngineServicer(), server)
    server.add_insecure_port('[::]:50051')
    print("Brain Service started on port 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()