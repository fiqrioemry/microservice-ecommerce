import React from "react";

const InputElement = ({
  label,
  name,
  error,
  register,
  placeholder,
  type = "text",
  style = "col-span-1 space-y-1",
}) => {
  return (
    <div className={style}>
      {label && (
        <label htmlFor={name} className="label">
          {label}
        </label>
      )}
      <input
        id={name}
        name={name}
        type={type}
        placeholder={placeholder}
        {...(register ? register(name) : {})}
        className={`input ${error ? "input-error" : ""}`}
      />
      {error && <p className="error-message">{error.message}</p>}
    </div>
  );
};

export default InputElement;
