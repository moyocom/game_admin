package moyoLogic

import (
	"fmt"
	"strings"
	"time"

	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/data_type"

	_ "lib/mysql"
)

func (this *TypLogDataBase) parseLine(serverId int, str string) *TypMap {
	argArr := strings.Split(str, "|")
	intTime := Int(argArr[0])
	retLine := HashMap(
		":time", intTime,
		":type", Int(argArr[1]))
	str_time := time.Unix(int64(intTime), 0).Format("2006-01-02 15:04:05")
	retLine.Assoc(":timeStr", str_time)
	typeId := retLine.Get(":type").(int)
	switch typeId {
	//玩家金钱增加.
	case 331:
		retLine.Assoc(
			":playerId", Int(argArr[2]),
			":amount", Int(argArr[3]),
			":moneyType", Int(argArr[4]),
			":payType", Int(argArr[5]))
		break
		//玩家获得物品
	case 324:
		retLine.Assoc(
			":playerId", argArr[2],
			":goodsType", argArr[3],
			":goodsId", argArr[4],
			":number", argArr[5],
			":getWay", argArr[6],
		)
		if len(argArr) > 7 {
			retLine.Assoc(":shopId", argArr[7])
		}
		break
		//玩家消耗物品
	case 325:
		retLine.Assoc(
			":playerId", argArr[2],
			":goodsType", argArr[3],
			":goodsId", argArr[4],
			":number", argArr[5],
			":way", argArr[6])

		break
		//玩家宠物出战/休息
	case 412, 411:
		retLine.Assoc(
			":playerId", argArr[2],
			":petId", argArr[3],
			":petType", argArr[4])
		break
		//玩家获取宠物
	case 418:
		retLine.Assoc(
			":playerId", argArr[2],
			":petId", argArr[3])
		break
		//玩家离开副本
	case 352:
		retLine.Assoc(
			":playerId", argArr[2],
			":copySceneId", argArr[3],
			":quitType", argArr[4],
			":monsterNumber", argArr[5])
		break
		//玩家进入副本
	case 351:
		retLine.Assoc(
			":playerId", argArr[2],
			":copySceneId", argArr[3],
			":enterType", argArr[4])
		break
		//玩家召唤队伍/解散队伍/加入队伍/创建队伍
	case 345, 344, 342, 341:
		retLine.Assoc(":playerId", argArr[2], ":teamId", argArr[3])
		break
		//玩家离开队伍
	case 343:
		retLine.Assoc(":playerId", argArr[2], ":teamId", argArr[3], ":quitType", argArr[4])
		break
		//玩家杀怪升级经验
	case 355:
		retLine.Assoc(":playerId", argArr[2], ":exp", argArr[3], ":monserId", argArr[4])
		break
		//玩家经验提升
	case 335:
		retLine.Assoc(":playerId", argArr[2], ":exp", argArr[3], ":type", argArr[4])
		break
		//玩家升级／激活技能
	case 323:
		retLine.Assoc(":playerId", argArr[2], ":skillId", argArr[3], ":skillLevel", argArr[4])
		break
		//玩家任务失败/取消任务/完成任务/达到任务完成条件
	case 315, 314, 313, 312:
		retLine.Assoc(":playerId", Int(argArr[2]), ":taskId", Int(argArr[3]), ":time2", argArr[4])
		break
		//玩家领取任务
	case 311:
		retLine.Assoc(":playerId", Int(argArr[2]), ":taskId", Int(argArr[3]))
		break
		//玩家复活
	case 308:
		retLine.Assoc(":playerId", argArr[2], ":type", argArr[3])
		break
		//玩家阅读邮件/删除邮件
	case 423, 422:
		retLine.Assoc(":playerId", argArr[2], ":mailId", argArr[3])
		break
		//发送邮件
	case 421:
		retLine.Assoc(
			":recvId", argArr[2],
			":senderId", argArr[3],
			":goodsId", argArr[4],
			":goodsNumber", argArr[5],
			":moneyType", argArr[6],
			":moneyNumber", argArr[7])
		break
		//玩家使用小喇叭
	case 306:
		retLine.Assoc(":playerId", argArr[2])
		break
		//玩家切换地图
	case 305:
		retLine.Assoc(
			":playerId", argArr[2],
			":curMapId", argArr[3],
			":targetMapId", argArr[4],
			":posX", argArr[5], ":posZ", argArr[6])
		break
		//切换PK模式
	case 304:
		retLine.Assoc(":playerId", argArr[2], ":pkMode", argArr[3])
		break
		//玩家死亡信息
	case 302:
		retLine.Assoc(
			":playerId", argArr[2], ":monsterId", argArr[3],
			":type", argArr[4], ":mapId", argArr[5],
			":posX", argArr[6], ":posZ", argArr[7], ":huodong", argArr[8])
		break
		//玩家死亡信息
	case 303:
		retLine.Assoc(":playerId", argArr[2], ":monsterId", argArr[3])
		break
		//玩家升级信息
	case 301:
		retLine.Assoc(":playerId", argArr[2])
		if len(argArr) > 3 {
			retLine.Assoc(":level", argArr[3])
			fmt.Println("have level?")
		}
		break
		//玩家创建重复名称
	case 205:
		retLine.Assoc(":name", argArr[2])
		break
		//玩家登出
	case 204:
		retLine.Assoc(":playerId", argArr[2])
		break
		//玩家登陆
	case 203:
		retLine.Assoc(":playerId", argArr[2], ":ip", argArr[3])
		break
		//角色创建
	case 201:
		retLine.Assoc(":playerId", argArr[2], ":name", argArr[3], ":accountId", argArr[4], ":sex", argArr[5])
		break
	}
	if this.LogData.Get(typeId) == nil {
		this.LogData.Assoc(typeId, List())
	}
	this.LogData.Get(typeId).(ICollection).Conj(retLine)
	return retLine
}
