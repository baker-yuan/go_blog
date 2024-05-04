import { FriendLink } from "@/api/interface";

// 友链状态 1-已发布 2-以下线
export const statusOptions = [
  {
    label: "已发布",
    value: FriendLink.FriendLinkStatus.STATUS_PUBLISHED
  },
  {
    label: "以下线",
    value: FriendLink.FriendLinkStatus.STATUS_OFFLINE
  },
  {
    label: "未知",
    value: FriendLink.FriendLinkStatus.STATUS_UNKNOWN
  }
];

// 友链状态 1-已发布 2-以下线
export const formatStatus = (status: number) => {
  if (status === FriendLink.FriendLinkStatus.STATUS_PUBLISHED) {
    return "已发布";
  } else if (status === FriendLink.FriendLinkStatus.STATUS_OFFLINE) {
    return "以下线";
  } else {
    return "未知";
  }
};
