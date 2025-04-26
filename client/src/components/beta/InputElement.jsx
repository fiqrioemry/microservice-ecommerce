// src/components/form/InputElement.jsx

import { Controller, useFormContext } from "react-hook-form";

const InputElement = ({
  name,
  label,
  rows = 4,
  placeholder = "",
  type = "text",
  isTextArea = false,
  disabled = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      rules={rules}
      render={({ field, fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label className="block text-sm font-medium">{label}</label>
          )}

          {isTextArea ? (
            <textarea
              {...field}
              placeholder={placeholder}
              disabled={disabled}
              rows={rows}
              className="w-full border  p-2 rounded resize-none disabled:bg-gray-100"
            />
          ) : (
            <input
              {...field}
              type={type}
              placeholder={placeholder}
              disabled={disabled}
              className="w-full border p-2 rounded disabled:bg-gray-100"
            />
          )}

          {fieldState.error && (
            <p className="text-red-500 text-xs mt-1">
              {fieldState.error.message || "This field is required"}
            </p>
          )}
        </div>
      )}
    />
  );
};

export { InputElement };
