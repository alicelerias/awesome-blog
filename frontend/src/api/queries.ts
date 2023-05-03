import { QueryFunctionContext } from "react-query";
import configs from "../configs/configs";
import {
  HealthCheck,
  Users,
  User,
  Posts,
  Post,
  Comments,
  Favorites,
  Favorite,
} from "../types";
import axios from "./axios";

export const getFavorites = async (): Promise<Posts> => {
  const url = new URL(configs.API_URL + "/favorites");

  const { data } = await axios.get<Posts>(url.toString());

  return data;
};

export const getFavoritesCount = async (
  id: string | null
): Promise<Favorites> => {
  const url = new URL(configs.API_URL + `/favorites/${id}`);

  const { data } = await axios.get<Favorites>(url.toString());

  return data;
};

export const isFollowing = async (id: string | null): Promise<Favorite> => {
  const url = new URL(configs.API_URL + `/follows/${id}`);

  const { data } = await axios.get<Favorite>(url.toString());

  return data;
};

export const getUsers = async (): Promise<Users> => {
  const url = new URL(configs.API_URL + "/users");

  const { data } = await axios.get<Users>(url.toString());

  return data;
};

export const getUser = async (id: string | null): Promise<User> => {
  const url = new URL(configs.API_URL + `/users/${id}`);

  const { data } = await axios.get<User>(url.toString());

  return data;
};

export const getCurrentUser = async (): Promise<User> => {
  const url = new URL(configs.API_URL + "/profile");

  const { data } = await axios.get<User>(url.toString());

  return data;
};

export const getAllPosts = async ({
  pageParam = "/posts",
}: QueryFunctionContext): Promise<Posts> => {
  const url = new URL(configs.API_URL + pageParam);

  const { data } = await axios.get<Posts>(url.toString());

  return data;
};

export const getPostsByUser = async ({
  pageParam = "/posts/you",
}: QueryFunctionContext): Promise<Posts> => {
  const url = new URL(configs.API_URL + pageParam);

  const { data } = await axios.get<Posts>(url.toString());

  return data;
};

export const getBlogsPost = async (id: string | null): Promise<Posts> => {
  const url = new URL(configs.API_URL + `/posts/blog/${id}`);

  const { data } = await axios.get<Posts>(url.toString());

  return data;
};

export const getFeed = async ({
  pageParam = "/feed",
}: QueryFunctionContext): Promise<Posts> => {
  const url = new URL(configs.API_URL + pageParam);

  const { data } = await axios.get<Posts>(url.toString());

  return data;
};

export const getPost = async (id: string | null): Promise<Post> => {
  const url = new URL(configs.API_URL + `/posts/${id}`);

  const { data } = await axios.get<Post>(url.toString());

  return data;
};

export const getComments = async (id: string | null): Promise<Comments> => {
  const url = new URL(configs.API_URL + `/comment/${id}`);

  const { data } = await axios.get<Comments>(url.toString());

  return data;
};

export const healthCheck = async (): Promise<HealthCheck> => {
  const url = new URL(configs.API_URL + "/healthcheck");

  const { data } = await axios.get<HealthCheck>(url.toString());

  return data;
};
