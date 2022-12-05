const displayTime = {
   prepareHoursForDisplay(hours) {
       return hours < 10 ? '0' + hours + ':00' : hours + ':00';
   }
};

export default displayTime;