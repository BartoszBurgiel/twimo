// Import core functionalities from react, react-nativ and expo
import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import color utility
import color from "../../utils/colorPallet";

// Define CircleIcon component
// Receives iconName and boolean status as properties from parent component
const CircleIcon = ({ iconName, isActive }) => {
  let activityColor = color.brand.light;
  if (isActive) {
    activityColor = color.brand.normal;
  }

  //FIXME: Make this thing more intutitive to use through generalization
  return (
    <View
      style={{
        backgroundColor: activityColor,
        height: 50,
        width: 50,
        margin: 5,
        borderRadius: 100,
        justifyContent: "center",
        alignItems: "center"
      }}
    >
      <Ionicons name={iconName} size={30} color={color.white} />
    </View>
  );
};

const SupplyIcons = ({ coffee, wifi, power, music, food }) => {
  return (
    <View style={styles.container}>
      <CircleIcon iconName="md-cafe" isActive={coffee ? true : false} />
      <CircleIcon iconName="md-wifi" isActive={wifi ? true : false} />
      <CircleIcon
        iconName="md-battery-charging"
        isActive={power ? true : false}
      />
      <CircleIcon iconName="md-musical-note" isActive={music ? true : false} />
      <CircleIcon iconName="md-pizza" isActive={food ? true : false} />
    </View>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flexWrap: "wrap",
    alignItems: "flex-start",
    flexDirection: "row",
    justifyContent: "space-between"
  }
});

// Export SupplyIcons to get used across the App
export default SupplyIcons;
