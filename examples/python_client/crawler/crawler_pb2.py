# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: crawler/crawler.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='crawler/crawler.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\036com.protoconf.examples.crawler'),
  serialized_pb=_b('\n\x15\x63rawler/crawler.proto\"M\n\x07\x43rawler\x12\x12\n\nuser_agent\x18\x01 \x01(\t\x12\x14\n\x0chttp_timeout\x18\x02 \x01(\x05\x12\x18\n\x10\x66ollow_redirects\x18\x03 \x01(\x08\"\xed\x01\n\x0e\x43rawlerService\x12\x1a\n\x08\x63rawlers\x18\x01 \x03(\x0b\x32\x08.Crawler\x12+\n\x06\x61\x64mins\x18\x02 \x03(\x0b\x32\x1b.CrawlerService.AdminsEntry\x12\x11\n\tlog_level\x18\x03 \x01(\x05\x1aN\n\x0b\x41\x64minsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12.\n\x05value\x18\x02 \x01(\x0e\x32\x1f.CrawlerService.AdminPermission:\x02\x38\x01\"/\n\x0f\x41\x64minPermission\x12\x0e\n\nREAD_WRITE\x10\x00\x12\x0c\n\x08GOD_MODE\x10\x01\x42 \n\x1e\x63om.protoconf.examples.crawlerb\x06proto3')
)



_CRAWLERSERVICE_ADMINPERMISSION = _descriptor.EnumDescriptor(
  name='AdminPermission',
  full_name='CrawlerService.AdminPermission',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='READ_WRITE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GOD_MODE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=295,
  serialized_end=342,
)
_sym_db.RegisterEnumDescriptor(_CRAWLERSERVICE_ADMINPERMISSION)


_CRAWLER = _descriptor.Descriptor(
  name='Crawler',
  full_name='Crawler',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_agent', full_name='Crawler.user_agent', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='http_timeout', full_name='Crawler.http_timeout', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='follow_redirects', full_name='Crawler.follow_redirects', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=102,
)


_CRAWLERSERVICE_ADMINSENTRY = _descriptor.Descriptor(
  name='AdminsEntry',
  full_name='CrawlerService.AdminsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='CrawlerService.AdminsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='CrawlerService.AdminsEntry.value', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=293,
)

_CRAWLERSERVICE = _descriptor.Descriptor(
  name='CrawlerService',
  full_name='CrawlerService',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='crawlers', full_name='CrawlerService.crawlers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='admins', full_name='CrawlerService.admins', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='log_level', full_name='CrawlerService.log_level', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_CRAWLERSERVICE_ADMINSENTRY, ],
  enum_types=[
    _CRAWLERSERVICE_ADMINPERMISSION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=105,
  serialized_end=342,
)

_CRAWLERSERVICE_ADMINSENTRY.fields_by_name['value'].enum_type = _CRAWLERSERVICE_ADMINPERMISSION
_CRAWLERSERVICE_ADMINSENTRY.containing_type = _CRAWLERSERVICE
_CRAWLERSERVICE.fields_by_name['crawlers'].message_type = _CRAWLER
_CRAWLERSERVICE.fields_by_name['admins'].message_type = _CRAWLERSERVICE_ADMINSENTRY
_CRAWLERSERVICE_ADMINPERMISSION.containing_type = _CRAWLERSERVICE
DESCRIPTOR.message_types_by_name['Crawler'] = _CRAWLER
DESCRIPTOR.message_types_by_name['CrawlerService'] = _CRAWLERSERVICE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Crawler = _reflection.GeneratedProtocolMessageType('Crawler', (_message.Message,), dict(
  DESCRIPTOR = _CRAWLER,
  __module__ = 'crawler.crawler_pb2'
  # @@protoc_insertion_point(class_scope:Crawler)
  ))
_sym_db.RegisterMessage(Crawler)

CrawlerService = _reflection.GeneratedProtocolMessageType('CrawlerService', (_message.Message,), dict(

  AdminsEntry = _reflection.GeneratedProtocolMessageType('AdminsEntry', (_message.Message,), dict(
    DESCRIPTOR = _CRAWLERSERVICE_ADMINSENTRY,
    __module__ = 'crawler.crawler_pb2'
    # @@protoc_insertion_point(class_scope:CrawlerService.AdminsEntry)
    ))
  ,
  DESCRIPTOR = _CRAWLERSERVICE,
  __module__ = 'crawler.crawler_pb2'
  # @@protoc_insertion_point(class_scope:CrawlerService)
  ))
_sym_db.RegisterMessage(CrawlerService)
_sym_db.RegisterMessage(CrawlerService.AdminsEntry)


DESCRIPTOR._options = None
_CRAWLERSERVICE_ADMINSENTRY._options = None
# @@protoc_insertion_point(module_scope)