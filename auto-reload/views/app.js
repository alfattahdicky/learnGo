
// console.log("hello")

setInterval(() => {
  window.location.reload();
}, 15000)

fetch("http://127.0.0.1:5174/", {
    method: "GET",
    header:{
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials" : true 
    }
  })
    .then(response => response.json())
    .then(result => {
      const {water, wind} = result.status
      textView("statusWater", conditionWater(water), "water")
      textView("statusWind", conditionWind(wind), "wind")
    })

const textView  = (idStatus, condition, idBackground) => {
  const statusEl = document.getElementById(idStatus);
  const backgroundStatus = document.getElementById(idBackground);
  switch (condition) {
    case "Aman":
      statusEl.textContent = "Aman";
      backgroundStatus.style.backgroundColor = "rgb(98, 238, 17)"
      break;
    case "Siaga":
      statusEl.textContent = "Siaga";
      backgroundStatus.style.backgroundColor = "#FF6D28";
      break;
    case "Bahaya":
      statusEl.textContent = "Bahaya";
      backgroundStatus.style.backgroundColor = "rgb(204, 19, 19)";
      break;
  }

} 

const conditionWater = (water) => {
  let status = "";

  if (water <= 5) {
    status += "Aman";
  }else if(water >= 6 && water <= 8) {
    status += "Siaga";
  }else if(water > 8) {
    status += "Bahaya";
  }

  return status;
}

const conditionWind = (wind) => {
  let status = "";
  switch (true) {
    case (wind <= 6):
      status += "Aman";
      break;
    case(wind >= 7 && wind <= 15):
      status += "Siaga";
      break;
    case(wind > 15):
      status += "Bahaya";
      break;
  }

  return status;
}


