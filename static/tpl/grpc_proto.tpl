syntax = "proto3";

package protos;
option go_package = "api/grpc/protos;protos";


service Greeting{
  	rpc GetList(GetGreetingListParam) returns (GetGreetingListReply) {}
  	rpc GetAll(GetGreetingAllParam) returns (GetGreetingAllReply) {}
  	rpc GetDetail(GetGreetingDetailParam) returns (GetGreetingDetailReply) {}
  	rpc Create(CreateGreetingParam) returns (CreateGreetingReply) {}
  	rpc Update(UpdateGreetingParam) returns (UpdateGreetingReply) {}
  	rpc Delete(DeleteGreetingParam) returns (DeleteGreetingReply) {}
}

message GetGreetingListParam{
  int64 pageIndex = 1;
  int64 pageSize = 2;

}
message GetGreetingListReply{
	message List {
	}
	repeated List item = 1;
	int64 total = 2;
}

message GetGreetingAllParam{
  
}

message GetGreetingAllReply{
	  message List {
	  }
	  repeated List item = 1;
}

message GetGreetingDetailParam{
  int64 Id  = 1;
}

message GetGreetingDetailReply{
}

message CreateGreetingParam{

}
  
message CreateGreetingReply{
  
}

message UpdateGreetingParam{
	int64 Id  = 1;
}
  
message UpdateGreetingReply{
  
}

message DeleteGreetingParam{
	int64 Id  = 1;
}
  
message DeleteGreetingReply{
  
}
