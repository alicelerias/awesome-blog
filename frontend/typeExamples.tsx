import React, { DetailedReactHTMLElement, ReactElement } from "react";

// export type Person = User & {
//   value: number;
//   name: string;
//   active: boolean;
// };

// export type GetPersonType = (id: number) => Person;

// export const getPerson:GetPersonType = (id) => {
//     return {} as Person;
// }

// export const getPerson = (id: number): Person => {
//   return {
//     value: 30,
//     name: "timtim",
//     active: true,
//   };
// };

// function getPersonB(id: number): Person {
//   return getPerson(id);
// }

// const Component = (props: { name: string}={name: "default"}): DetailedReactHTMLElement<P, T> => {
//     return React.createElement;
// }

type props = {
  name: string  
}
const Component: React.FC<props> = ({name}) => {
  return (
    <h1>{name}</h1>
    // <h1></h1>
  );
};

// type User = {name: string}
// type genericType<P extends User> = (user: P, context : React.Context) => ReactElement

// type example = genericType<Person>