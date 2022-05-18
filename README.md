This is a homework duplicate of the demo project [**easy_note**](https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note) from kitex project.

Compared with the original project, the major difference is the IDL part. I implemented
the **note** part with ProtoBuf as the **user**. And I found that the default generated
template code in golang is different for **string** type. Thrift always generate string
type as pointer of string while ProtoBuf as string. It may differ in performance for
treatment of long strings.

And the CRUD functions in DAL part has a misunderstanding behavior. For [dal/db/user.go](https://github.com/cloudwego/kitex-examples/blob/main/bizdemo/easy_note/cmd/user/dal/db/user.go),
the function MGetUsers return nil under the condition that database operation failure, while
the [dal/db/note.go](https://github.com/cloudwego/kitex-examples/blob/main/bizdemo/easy_note/cmd/note/dal/db/note.go)
return an empty array.