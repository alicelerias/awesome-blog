import configs from "./../configs/configs";
import { Credential, Post, User, Comment } from "../types";
import axios from "./axios";

export const registerUser = async (input: User) => {
  const url = new URL(configs.API_URL + "/register");
  await axios.post(url.toString(), input);
};

export const login = async (input: Credential) => {
  const url = new URL(configs.API_URL + "/login");
  await axios.post(url.toString(), input);
};

export const logout = async () => {
  await axios.post(configs.API_URL + "/logout");
};

export const favorite = async (id: string | null) => {
  const url = new URL(configs.API_URL + `/favorite/${id}`);
  await axios.post(url.toString());
};

export const unfavorite = async (id: string | null) => {
  const url = new URL(configs.API_URL + `/favorite/${id}`);
  await axios.delete(url.toString());
};

export const Follow = async (id: string | undefined) => {
  const url = new URL(configs.API_URL + `/follow/${id}`);
  await axios.post(url.toString());
};

export const unfollow = async (id: string | undefined) => {
  const url = new URL(configs.API_URL + `/follow/${id}`);
  await axios.delete(url.toString());
};

export const createComment = async (id: string | null, input: Comment) => {
  const url = new URL(configs.API_URL + `/comment/${id}`);
  await axios.post(url.toString(), input);
};

export const deleteComment = async (id: string | null) => {
  const url = new URL(configs.API_URL + `/comment/${id}`);
  await axios.delete(url.toString());
};

export const updateCurrentUser = async (input: User) => {
  const url = new URL(configs.API_URL + "/profile");
  await axios.put(url.toString(), input);
};

export const updatePost = async (id: string | null, input: Post) => {
  const url = new URL(configs.API_URL + `/posts/${id}`);
  await axios.put(url.toString(), input);
};

export const createPost = async (input: Post) => {
  const url = new URL(configs.API_URL + "/posts");
  await axios.post(url.toString(), input);
};

export const deletePost = async (id: string | null) => {
  const url = new URL(configs.API_URL + `/posts/${id}`);
  await axios.delete(url.toString());
};
