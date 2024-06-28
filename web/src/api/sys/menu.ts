import { defHttp } from '@/utils/http/axios';
import { getMenuListResultModel } from './model/menuModel';

enum Api {
  GetMenuList = '/v1/menus/meta',
  CreateMenu = '/v1/menus',
}

/**
 * @description: Get user menu based on id
 */

export const getMenuList = () => {
  return defHttp.get<getMenuListResultModel>({ url: Api.GetMenuList });
};

export const createMenu = (data) => {
  return defHttp.post({ url: Api.CreateMenu, data });
};
