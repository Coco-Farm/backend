const HexColors = {
    red: "#FF0000",     // 빨강
    orange: "#FFA500",  // 주황
    yellow: "#FFFF00",  // 노랑
    green: "#008000",   // 초록
    blue: "#0000FF",    // 파랑
    indigo: "#4B0082",  // 남색
    violet: "#EE82EE"   // 보라
};

// HexColors 객체의 키(색상)를 배열로 가져오기
const colorsArray = Object.keys(HexColors);

// 랜덤하게 색상 선택 함수
function getRandomColor() {
    const randomIndex = Math.floor(Math.random() * colorsArray.length);
    return HexColors[colorsArray[randomIndex]];
}

// 테스트: 랜덤 색상 선택 및 출력
const randomColor = getRandomColor();
console.log("랜덤 색상:", randomColor);
