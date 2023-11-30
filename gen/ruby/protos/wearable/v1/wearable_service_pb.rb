# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/wearable/v1/wearable_service.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("protos/wearable/v1/wearable_service.proto", :syntax => :proto3) do
    add_message "wearable.v1.BeatsPerMinuteRequest" do
      optional :uuid, :string, 1
    end
    add_message "wearable.v1.BeatsPerMinuteResponse" do
      optional :value, :uint32, 1
      optional :minute, :uint32, 2
    end
  end
end

module Wearable
  module V1
    BeatsPerMinuteRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("wearable.v1.BeatsPerMinuteRequest").msgclass
    BeatsPerMinuteResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("wearable.v1.BeatsPerMinuteResponse").msgclass
  end
end