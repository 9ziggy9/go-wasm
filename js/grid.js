function Grid (gridObj) {
  const {cols,rows,scale,containerId,gridId,cellClass} = gridObj;

  // PARAMETER CHECKING
  const params = ["cols","rows","scale","containerId","gridId","cellClass"];
  const givenParams = Object.keys(gridObj);
  const missingParams = params.filter(k => !givenParams.includes(k));
  console.assert(
    !missingParams.length,
    `Missing parameters: ${missingParams}`
  );

  const container = document.getElementById(containerId);
  const grid = document.getElementById(gridId);
  const aspectRatio = cols / rows;
  container.style.width = `${aspectRatio * scale}vmin`;
  container.style.height = `${scale}vmin`;
  grid.style.display = "grid";
  grid.style.width = "100%";
  grid.style.height = "100%";
  grid.style.gridTemplateColumns = `repeat(${cols}, 1fr)`;
  grid.style.gridTemplateRows = `repeat(${rows}, 1fr)`;
  return {
    "destroy": () => grid.innerHTML = "",
    "create": () => {
      for (let y = 0; y < rows; y++) {
	for (let x = 0; x < cols; x++) {
	  const cell = document.createElement("div");
	  cell.classList.add(cellClass);
	  cell.setAttribute("id", `(${x},${y})`);
	  grid.appendChild(cell);
	}
      }
    },
    "bindListeners": (event, ...bindings) => {
      for (const child of grid.children) {
	bindings.forEach((b) => child.addEventListener(event, () => b(child)));
      }
    },
  };
}
