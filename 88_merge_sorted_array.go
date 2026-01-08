//Para resolver esse problema iteramos sobre os arrays de tras para frente, com dois indices(um para cada array)
// e um terceiro index, para sabermos o ultimo valor valido no array 1 (que é sempre maior)
// toda iteração deve adicionar um valor ao indice atual de nums1( o index que começou no final) e fazer pelo
//menos uma alteração  no array de nums1, SEMPRE COM O MAIOR VALOR PRESENTE NO LOOP (vindo ou de nums1 ou nums2)

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1 // Último elemento válido em nums1
	p2 := n - 1 // Último elemento em nums2
	p := m + n - 1 // Última posição total de nums1
	for p2 >= 0 {
		// Se ainda houver elementos em nums1 e o valor em nums1 for maior
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			// Caso contrário, ou se nums1 acabou, pega de nums2
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

}