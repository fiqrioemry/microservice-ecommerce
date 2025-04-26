import { Controller, useFormContext } from "react-hook-form";

const InputElement = ({
  name,
  label,
  rows = 4,
  type = "text",
  placeholder = "",
  isNumber = false,
  disabled = false,
  isTextArea = false,
  rules = { required: true },
  maxLength, // ➡️ Tambahan baru
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleKeyDown = (e) => {
          if (isNumber) {
            if (
              [
                "Backspace",
                "Tab",
                "Delete",
                "ArrowLeft",
                "ArrowRight",
              ].includes(e.key)
            ) {
              return;
            }
            if (!/^[0-9]$/.test(e.key)) {
              e.preventDefault();
            }
          }
        };

        return (
          <div className="space-y-1">
            {label && (
              <label className="block text-sm font-medium">{label}</label>
            )}

            {isTextArea ? (
              <textarea
                {...field}
                value={field.value ?? ""}
                placeholder={placeholder}
                disabled={disabled}
                rows={rows}
                maxLength={maxLength || undefined}
                className="w-full border p-2 rounded resize-none disabled:bg-gray-100"
              />
            ) : (
              <input
                {...field}
                value={field.value ?? ""}
                type={isNumber ? "text" : type}
                onKeyDown={handleKeyDown}
                placeholder={placeholder}
                disabled={disabled}
                inputMode={isNumber ? "numeric" : undefined}
                maxLength={maxLength || undefined}
                className="w-full border p-2 rounded disabled:bg-gray-100"
              />
            )}

            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputElement };
