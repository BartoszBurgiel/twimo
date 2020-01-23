// Import core functionality dependencies
import React from "react";
import { View, Text, StyleSheet, Image } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Import custom built components
import RatingStars from "../components/RatingStars";
import SupplyIcons from "../components/SupplyIcons";
import ReviewCard from "../components/ReviewCard";

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
      <View style={styles.contentWrapper}>
        <Text style={styles.categoryTag}>coffeeshop</Text>
        <Text style={styles.headerText}>Starbucks Montabaur</Text>
        <RatingStars colorfill={color.brand.normal} rating={4} />
        <SupplyIcons />
        <View style={styles.horizontalRule} />
        <Text style={styles.secondaryHeaderText}>Reviews</Text>
        <ReviewCard />
      </View>
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
  contentWrapper: {
    paddingHorizontal: 30
  },
  categoryTag: {
    fontSize: 20,
    fontWeight: "300",
    textTransform: "uppercase"
  },
  headerText: {
    fontSize: 30,
    fontWeight: "700"
  },
  horizontalRule: {
    borderBottomColor: color.gray.dark,
    borderBottomWidth: 1,
    marginVertical: 20
  },
  secondaryHeaderText: {
    fontSize: 20,
    fontWeight: "500"
  }
});

// Export the above defined component
export default LocationDetailsScreen;
