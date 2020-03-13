function dataHandle(str) {
  let arr = JSON.parse(str);
  arr = arr.filter(item => {
    return parent(item.iid) === 0;
  });
  return arr;
}

module.exports = { dataHandle };
