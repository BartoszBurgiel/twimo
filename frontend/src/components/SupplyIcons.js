// Import core functionalities
import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import color utility
import color from "../../utils/colorPallet";

// Define the circular icons
const CircleIcon = ({ iconName, isTrue }) => {
  let activityColor = color.brand.light;
  if (isTrue) {
    activityColor = color.brand.normal;
  }

  //FIXME: Center Icon in the circular view prop
  return (
    <View
      style={{
        backgroundColor: activityColor,
        height: 50,
        width: 50,
        margin: 5,
        borderRadius: 100,
        flexDirection: "row",
        alignItems: "center",
        paddingHorizontal: 10
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

const styles = StyleSheet.create({
  container: {
    flexWrap: "wrap",
    alignItems: "flex-start",
    flexDirection: "row"
  },
  icon: {
    margin: 10
  }
});

export default SupplyIcons;
