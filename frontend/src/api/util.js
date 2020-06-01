function getURL(base) {
  return process.env.NODE_ENV === 'development'
    ? `http://localhost:3000${base}`
    : base;
}

export {
  getURL,
};
