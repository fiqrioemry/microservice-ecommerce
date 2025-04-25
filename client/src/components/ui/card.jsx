const Card = ({ children, className = "" }) => {
  return <div className={`card card-hover ${className}`}>{children}</div>;
};

export default Card;
