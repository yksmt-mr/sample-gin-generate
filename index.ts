import { UserApi } from "./api/axios/index";

const api = new UserApi();

const res = await api.getUsersUserId(1).catch((e) => console.log(e));

if (res) {
  console.log(res.data.lastName);
}
