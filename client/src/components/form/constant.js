export const loginFormSchema = [
  {
    name: "email",
    label: "Email",
    type: "input",
    inputType: "email",
    placeholder: "Enter your email",
    required: true,
  },
  {
    name: "password",
    label: "Password",
    type: "input",
    inputType: "password",
    placeholder: "Enter your password",
    required: true,
  },
  {
    name: "gender",
    label: "Gender",
    type: "select",
    isAsync: false,
    options: [
      { label: "Male", value: "male" },
      { label: "Female", value: "female" },
    ],
    placeholder: "Select gender",
    required: true,
  },
  {
    name: "province",
    label: "Province",
    type: "select",
    isAsync: true,
    loaderKey: "province",
    placeholder: "Select province",
    required: true,
  },
];
