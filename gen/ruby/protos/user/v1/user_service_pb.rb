# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/user/v1/user_service.proto

require 'google/protobuf'

require 'protos/user/v1/user_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("protos/user/v1/user_service.proto", :syntax => :proto3) do
    add_message "user.v1.GetUserRequest" do
      optional :uuid, :string, 1
    end
    add_message "user.v1.GetUserResponse" do
      optional :user, :message, 1, "user.v1.User"
    end
  end
end

module User
  module V1
    GetUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.v1.GetUserRequest").msgclass
    GetUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.v1.GetUserResponse").msgclass
  end
end
