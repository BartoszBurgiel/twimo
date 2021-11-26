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

// Define the FeatureIcons component
// Receives featuresObject as props from parent component
const FeatureIcons = ({ features }) => {
  return (
    <View style={styles.container}>
      <CircleIcon
        iconName="md-cafe"
        isActive={features.Coffee ? true : false}
      />
      <CircleIcon iconName="md-wifi" isActive={features.Wifi ? true : false} />
      <CircleIcon
        iconName="md-battery-charging"
        isActive={features.Power ? true : false}
      />
      <CircleIcon
        iconName="md-musical-note"
        isActive={features.Music ? true : false}
      />
      <CircleIcon iconName="md-pizza" isActive={features.Food ? true : false} />
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

// Export FeatureIcons to get used across the App
export default FeatureIcons;
