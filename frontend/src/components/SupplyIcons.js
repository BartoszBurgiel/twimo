// Import core functionalities from react, react-nativ and expo
import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import color utility
import color from "../../utils/colorPallet";

// Define CircleIcon component
// Receives iconName and boolean status as properties from parent component
const CircleIcon = ({ iconName, isTrue }) => {
  let activityColor = color.brand.light;
  if (isTrue) {
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

const SupplyIcons = () => {
  return (
    <View style={styles.container}>
      <CircleIcon iconName="md-cafe" isTrue />
      <CircleIcon iconName="md-wifi" isTrue />
      <CircleIcon iconName="md-battery-charging" isTrue />
      <CircleIcon iconName="md-musical-note" />
      <CircleIcon iconName="md-pizza" isTrue />
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
