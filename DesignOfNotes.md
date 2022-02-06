## Demand  of Note struct

Note 备忘录
is used to help people thinking about what to do.



## field

- model 自定义的 `gorm.Model`，`CreatedAt` 等三个时间字段改成 `int`
- event title
- member id
- event type 用枚举
  - birthday
  - working
  - homework
  - meeting
- status 事件状态
  - doing 正在做，时间还没到，事情还没做完
  - finished 在时间内已经做完了
  - time out 超时了
  - delay 超时后延迟结束时间，如果在新的结束时间内完成会变成 `finished`
- start time 事件开始时间
- end time 事件结束时间
- content 事件内容
- view 事件查看次数

## API

- check event
- get event
    - add view 每次查看事件增加事件查看次数
    - check status 检查事件是否已到结束时间
- delay event 推迟结束时间，无论何时调用这个函数，事件状态都会变成 delay
- delete event 删除事件
- create event 新建事件
- update event 更新事件信息

