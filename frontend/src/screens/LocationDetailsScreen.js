// Import core functionality dependencies
import React from "react";
import { View, Text, StyleSheet, Image } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

import RatingStars from "../components/RatingStars";
import SupplyIcons from "../components/SupplyIcons";

// Define Component
const LocationDetailsScreen = () => {
  return (
    <View style={styles.container}>
      <View style={styles.imgWrapper}>
        <Image
          style={styles.image}
          source={{
            uri:
              "https://www.heikaus-asset.com/wp-content/uploads/2018/01/Test_Starbucks-1024x731.jpg"
          }}
        ></Image>
      </View>
      <View style={styles.textWrapper}>
        <Text style={styles.category}>coffeeshop</Text>
        <Text style={styles.headerText}>Starbucks Montabaur</Text>
        <RatingStars colorfill={color.brand.normal} rating={4} />
        <SupplyIcons />
      </View>
      <Text>Separator</Text>
      <Text>UserRatings</Text>
    </View>
  );
};

// Define styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white
  },
  imgWrapper: {
    alignItems: "center",
    marginBottom: 20
  },
  image: {
    height: 275,
    width: 500
  },
  textWrapper: {
    paddingHorizontal: 30
  },
  category: {
    fontSize: 20,
    fontWeight: "300",
    textTransform: "uppercase"
  },
  headerText: {
    fontSize: 30,
    fontWeight: "900"
  }
});

// Export the above defined component
export default LocationDetailsScreen;
