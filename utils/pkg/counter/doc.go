package counter

// 计数器
// 功能: 通过缓存来计数，并指定时间定时重置缓存的计数。
// 场景:
// 1. 接收图片数量，凌晨重置为 0，可以知晓每天接收的图片数量。
// 2. 负载均衡，计算某个节点的负载情况，定时与实际负载进行同步。
